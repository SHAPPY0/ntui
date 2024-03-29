package nomad

import (
	// "fmt"
	"time"
	"sort"
	"github.com/hashicorp/nomad/api"
	"github.com/shappy0/ntui/internal/models"
)

type AllocationClient interface {
	List(*api.QueryOptions) ([]*api.AllocationListStub, *api.QueryMeta, error)
	Info(string, *api.QueryOptions) (*api.Allocation, *api.QueryMeta, error)
	Stats(*api.Allocation, *api.QueryOptions) (*api.AllocResourceUsage, error)
	Restart(*api.Allocation, string, *api.QueryOptions) error
}

func (n *Nomad) Allocations(allocName string, params *models.NomadParams) ([]models.Allocations, error) {
	Result := make([]models.Allocations, 0)
	Data, _, Err := n.AllocationClient.List(&api.QueryOptions{
		Namespace:	params.Namespace,
		Region:	params.Region,
	})
	if Err != nil {
		return Result, Err
	}
	for _, allocation := range Data {
		if allocation.TaskGroup == allocName {
			Task, _ := n.AllocationTask(params, allocation.ID)
			Events, AllocStatus := n.AllocationEvents(allocation)
			alloc := models.Allocations{
				ID:			allocation.ID,
				Name:		allocation.Name,
				TaskName:	Task[0].Name,
				JobID:		allocation.JobID,
				Namespace:	allocation.Namespace,
				TaskGroup:	allocation.TaskGroup,
				Created:	time.Unix(0, allocation.CreateTime),
				Modified:	time.Unix(0, allocation.ModifyTime),
				Status:		AllocStatus,
				Version:	int(allocation.JobVersion),
				Client:		allocation.NodeID,
				Tasks:		Task[0],
				Events:		Events,
				Volumn:		"N/A",
				Cpu:		Task[0].Resources.CPU,
				CpuUsage:	Task[0].Resources.CPUUsage,
				Memory:		Task[0].Resources.MemoryMB,
				MemoryUsage:Task[0].Resources.MemoryUsage,
			}
			Result = append(Result, alloc)
		}
	}
	return Result, nil
}

func (n *Nomad) AllocationTask(params *models.NomadParams, allocationId string) ([]models.Tasks, error) {
	Result := make([]models.Tasks, 0)

	Data, _, Err := n.AllocationClient.Info(allocationId, &api.QueryOptions{
		Namespace:	params.Namespace,
		Region:		params.Region,
	})
	if Err != nil {
		return Result, Err
	}
	for _, task := range Data.GetTaskGroup().Tasks{
		t := models.Tasks{
			Name:		task.Name,
			Driver:		task.Driver,
			Config:		task.Config,
			Env:		task.Env,
			Resources:	models.TaskResource{
				CPU:		*task.Resources.CPU,
				Cores:		*task.Resources.Cores,
				MemoryMB:	*task.Resources.MemoryMB,
				MemoryMaxMB: *task.Resources.MemoryMaxMB,
				DiskMB:		*task.Resources.DiskMB,
			},
			RestartPolicy:	models.TaskRestartPolicy{
				Interval:	*task.RestartPolicy.Interval,
				Attempts:	*task.RestartPolicy.Attempts,
				Delay:		*task.RestartPolicy.Attempts,
				Mode:		*task.RestartPolicy.Mode,
			},
		}
		Stats, _ := n.AllocationStats(params, Data)
		t.Resources.CPUUsage = int(Stats.Cpu.TotalTicks)
		t.Resources.MemoryUsage = int(Stats.Memory.RSS)
		Result = append(Result, t)
	}
	return Result, nil
}

func (n *Nomad) AllocationEvents(allocation *api.AllocationListStub) ([]models.Events, string) {
	Result := make([]models.Events, 0)
	AllocStatus := ""
	for _, taskStates := range allocation.TaskStates{
		AllocStatus = taskStates.State
		for _, event := range taskStates.Events {
			NewEvent := models.Events{
				Type:			event.Type,
				Time:			time.Unix(0, event.Time),
				DisplayMessage:	event.DisplayMessage,
				Message:		event.Message,
				FailsTask:		event.FailsTask,
				KillReason:		event.KillReason,
				KillTimeout:	event.KillTimeout,
				KillError:		event.KillError,
				DownloadError:	event.DownloadError,
				ValidationError: event.ValidationError,
				VaultError:		event.VaultError,
			}
			Result = append(Result, NewEvent)
		}
	}
	sort.Slice(Result, func (I, J int) bool {
		return Result[I].Time.After(Result[J].Time)
	})
	return Result, AllocStatus
}

func (n *Nomad) AllocationStats(params *models.NomadParams, allocation *api.Allocation) (models.AllocationStats, error) {
	Data, Err := n.AllocationClient.Stats(allocation, &api.QueryOptions{
		Namespace:	params.Namespace,
		Region:		params.Region,
	})
	if Err != nil {
		return models.AllocationStats{}, Err
	}
	Result := models.AllocationStats{
		Memory:		models.AllocMemoryStats{
			RSS:			Data.ResourceUsage.MemoryStats.RSS,
			Cache:			Data.ResourceUsage.MemoryStats.Cache,
			Swap:			Data.ResourceUsage.MemoryStats.Swap,
			Usage:			Data.ResourceUsage.MemoryStats.Usage,
			MaxUsage:		Data.ResourceUsage.MemoryStats.MaxUsage,
			KernelUsage:	Data.ResourceUsage.MemoryStats.KernelUsage,
			KernelMaxUsage:	Data.ResourceUsage.MemoryStats.KernelMaxUsage,
		},
		Cpu:	models.AllocCpuStats{
			SystemMode:			Data.ResourceUsage.CpuStats.SystemMode,
			UserMode:			Data.ResourceUsage.CpuStats.UserMode,
			TotalTicks:			Data.ResourceUsage.CpuStats.TotalTicks,
			ThrottledPeriods:	Data.ResourceUsage.CpuStats.ThrottledPeriods,
			ThrottledTime:		Data.ResourceUsage.CpuStats.ThrottledTime,
			Percent:			Data.ResourceUsage.CpuStats.Percent,
		},
	}
	return Result, nil
}

func (n *Nomad) Restart(allocId, taskName string, params *models.NomadParams) error {
	if Err := n.AllocationClient.Restart(&api.Allocation{ID: allocId}, taskName, &api.QueryOptions{
		Namespace:	params.Namespace,
		Region:		params.Region,
	}); Err != nil {
		return Err
	}
	return nil
}