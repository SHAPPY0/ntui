package nomad

import (
	"fmt"
	"time"
	"sort"
	"github.com/hashicorp/nomad/api"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

type JobClient interface {
	List(*api.QueryOptions) ([]*api.JobListStub, *api.QueryMeta, error) 
	Info(string, *api.QueryOptions) (*api.Job, *api.QueryMeta, error)
	Summary(string, *api.QueryOptions) (*api.JobSummary, *api.QueryMeta, error)
	Allocations(string, bool, *api.QueryOptions) ([]*api.AllocationListStub, *api.QueryMeta, error)
	Versions(string, bool, *api.QueryOptions) ([]*api.Job, []*api.JobDiff, *api.QueryMeta, error)
	Submission(string, int, *api.QueryOptions) (*api.JobSubmission, *api.QueryMeta, error)
	Revert(string, uint64, *uint64, *api.WriteOptions, string, string) (*api.JobRegisterResponse, *api.WriteMeta, error)
	Deregister(string, bool, *api.WriteOptions) (string, *api.WriteMeta, error)
	Register(*api.Job, *api.WriteOptions) (*api.JobRegisterResponse, *api.WriteMeta, error)
}

//Get jobs list
func (n *Nomad) Jobs(params *models.NomadParams) ([]models.Jobs, error) {
	Result := make([]models.Jobs, 0)
	Data, _, Err := n.JobClient.List(&api.QueryOptions{
		Region: params.Region,
		Namespace: params.Namespace,
	})
	if Err != nil {
		fmt.Println("Err", Err)
		return Result, Err
	}
	for Index, Job := range Data {
		_ = Index
		Summary := models.Summary{
			Total:	len(Job.JobSummary.Summary),
		}
		for _, JS := range Job.JobSummary.Summary{
			if JS.Running > 0 {
				Summary.Running++
			}
		}
		var J models.Jobs
		J.ID 			=	Job.ID
		J.Name			=	Job.Name
		J.Namespace		=	Job.JobSummary.Namespace
		J.Type			=	Job.Type
		J.Status		=	Job.Status
		J.StatusDescription = Job.StatusDescription
		J.Priority		=	Job.Priority
		J.StatusSummary = 	Summary
		J.SubmitTime	=	time.Unix(0, Job.SubmitTime)

		Result = append(Result, J)
	}
	return Result, nil
}

//Get TaskGroup list by jobId
func (n *Nomad) TaskGroups(jobId, region, namespace string) ([]models.TaskGroups, error) {
	Result := make([]models.TaskGroups, 0)
	Data, _, Err := n.JobClient.Summary(jobId, &api.QueryOptions{
		Region:		region,
		Namespace:  namespace,
	})
	if Err != nil {
		return Result, Err
	}

	for Key, Tg := range Data.Summary {
		TaskGroup := models.TaskGroups{
			Name:		Key,
			JobID:		Data.JobID,
			Queued:		Tg.Queued,
			Complete:	Tg.Complete,
			Failed:		Tg.Failed,
			Running:	Tg.Running,
			Starting:	Tg.Starting,
			Lost:		Tg.Lost,
			Unknown:	Tg.Unknown,
		} 
		Result = append(Result, TaskGroup)
	}
	sort.Slice(Result, func(I, J int) bool {
		return Result[I].Name < Result[J].Name
	})
	return Result, nil
}

//TODO: Implement Later, Submission(Definitions) required current job version
func (n *Nomad) Submission() {}

//Get versions List by jobId
func (n *Nomad) Versions(jobId string, params *models.NomadParams) ([]models.JobVersion, []models.JobVersionDiff, error) {
	JobVersions := make([]models.JobVersion, 0)
	JobVersionDiff := make([]models.JobVersionDiff, 0)
	Versions, Diffs, _, Err := n.JobClient.Versions(jobId, true, &api.QueryOptions {
		Region:		params.Region,
		Namespace:	params.Namespace,
	})
	if Err != nil {
		return JobVersions, JobVersionDiff, Err
	}
	for I := 0; I < len(Versions); I++{
		var Jv models.JobVersion
		Jv.Region		=	utils.SafeDeref(Versions[I].Region)
		Jv.Namespace	=	utils.SafeDeref(Versions[I].Namespace)
		Jv.ID			=	utils.SafeDeref(Versions[I].ID)
		Jv.Name			=	utils.SafeDeref(Versions[I].Name)
		Jv.Type			=	utils.SafeDeref(Versions[I].Type)
		Jv.Stop			=	utils.SafeDeref(Versions[I].Stop)
		Jv.Status		=	utils.SafeDeref(Versions[I].Status)
		Jv.Stable		=	utils.SafeDeref(Versions[I].Stable)
		Jv.Version		=	utils.SafeDeref(Versions[I].Version)
		Jv.SubmitTime	=	time.Unix(0, utils.SafeDeref(Versions[I].SubmitTime))
		JobVersions = append(JobVersions, Jv)
	}
	for I := 0; I < len(Diffs); I++ {
		var Jvd models.JobVersionDiff
		Jvd.Type 		= Diffs[I].Type
		Jvd.ID 			= Diffs[I].ID
		Jvd.Fields 		= Diffs[I].Fields
		Jvd.Objects 	= Diffs[I].Objects
		Jvd.TaskGroups 	= Diffs[I].TaskGroups
		JobVersionDiff = append(JobVersionDiff, Jvd)
	}
	return JobVersions, JobVersionDiff, nil 
}

//Revert job by version
func (n *Nomad) Revert(jobId string, version uint64, params *models.NomadParams) (error) {
	revertResp, _, err := n.JobClient.Revert(jobId, version, nil, nil, "", "")
	if err != nil {
		n.Logger.Error("Version " + jobId + " #" + utils.IntToStr(int(version)) + " revert failed  " + err.Error())
	}
	n.Logger.Info("Version " + jobId + " #" + utils.IntToStr(int(version)) + " reverted successful  " + utils.Stringify(revertResp))
	return err
}

//Deregister job
func (n *Nomad) Deregister(jobId string, purge bool, params *models.NomadParams) error {
	if jobId == "" {
		n.Logger.Error("Invalid jobId to stop job")
		return fmt.Errorf("Invalid jobId to stop job")
	}
	resp, _, err := n.JobClient.Deregister(jobId, purge, nil)
	if err != nil {
		n.Logger.Error("Job " + jobId + "stop failed: " + err.Error())
		return err
	}
	n.Logger.Info("Job " + jobId + " stopped successfully: " + resp)
	return nil

}

func (n *Nomad) GetJob(jobId string) (*api.Job, error) {
	job, _, err := n.JobClient.Info(jobId, nil)
	if err != nil {
		return nil, err
	}
	return job, nil
}

//Register Job
func (n *Nomad) Register(jobId string, params *models.NomadParams) error {
	if jobId == "" {
		n.Logger.Error("Invalid jobId to start job")
		return fmt.Errorf("Invalid jobId to start job")
	}
	job, err := n.GetJob(jobId)
	if err != nil {
		n.Logger.Error("Error getting job " + jobId + " info to start job: " + err.Error())
		return err
	}
	n.Logger.Info(utils.Stringify(&api.Job{Name: &jobId, Region: &params.Region, Namespace: &params.Namespace}))
	resp, _, err := n.JobClient.Register(job, nil)
	if err != nil {
		n.Logger.Error("Job " + jobId + "start failed: " + err.Error())
		return err
	}
	n.Logger.Info("Job " + jobId + " started successfully: " + utils.Stringify(resp))
	return nil
}