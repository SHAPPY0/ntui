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
	Deregister(string, bool, *api.WriteOptions) (string, *api.WriteMeta, error)
	Register(*api.Job, *api.WriteOptions) (*api.JobRegisterResponse, *api.WriteMeta, error)
}

func (n *Nomad) Jobs(params *models.NomadParams) ([]models.Jobs, error) {
	Result := make([]models.Jobs, 0)
	Data, _, Err := n.JobClient.List(&api.QueryOptions{
		Region: params.Region,
		Namespace: params.Namesapce,
	})
	if Err != nil {
		fmt.Println("Err", Err)
		return Result, Err
	}
	for Index, Job := range Data {
		_ = Index
		Summary := models.Summary{
			Total:	len(Job.JobSummary.Summary)
		}
		for _, JS := range Job.JobSummary.Summary{
			if JS.Running > 0 {
				Summary.Running++
			}
		}
		var J models.Jobs
		J.ID 			=	Job.ID
		J.Name			=	Job.Name
		J.Namespace		=	Job.JobSummary.Namesapce
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

func (n *Nomad) Versions(jobId string, params *models.NomadParams) ([]models.JobVersions, []models.JobVersionDiff, error) {
	JobVersions := make([]models.JobVersion, 0)
	JobVersionDiff := make([]models.JobVersionDiff, 0)
	Versions, Diff, _, Err := n.JobClient.Versions(jobId, true, &api.QueryOptions {
		Region:		params.Region,
		Namespace:	params.Namesapce,
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