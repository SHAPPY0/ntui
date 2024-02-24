package models

import (
	"time"
	"github.com/hashicorp/nomad/api"
)

type AlertMessage struct {
	Type 	string
	Text 	string
}

type AlertChan chan AlertMessage

type Regions struct {
	Id 		string
	Name 	string
}

type Namespaces struct {
	Id			string
	Name		string
	Description string
}

type NomadParams struct {
	Region 		string
	Namespace	string
	Datacenter  string
	JobId		string
}

type Jobs struct {
	ID					string
	Name				string
	Namespace			string
	Type 				string
	Status 				string
	StatusDescription 	string
	StatusSummary 		Summary
	Priority			int
	SubmitTime			time.Time
}

type Summary struct {
	Total 		int
	Running 	int
}

type TaskGroups struct {
	Name 			string
	JobID 			string
	Queued 			int
	Complete 		int
	Failed 			int
	Running 		int
	Starting 		int
	Lost 			int
	Unknown 		int
}

type Allocations struct {
	ID 				string
	Name 			string
	TaskName 		string
	JobID			string
	Namespace 		string
	TaskGroup 		string
	Created			time.Time
	Modified 		time.Time
	Status			string
	Version			int
	Client 			string
	Volumn 			string
	Cpu 			int
	CpuUsage 		int
	Memory 			int
	MemoryUsage 	int
	Tasks 			Tasks
	Events			[]Events
}

type TaskConfig struct {
	Image 			string
	MemoryHardLimit int
	Volumes 		[]string
}

type TaskResource struct {
	CPU 			int
	CPUUsage		int
	Cores			int
	MemoryMB		int
	MemoryUsage		int
	MemoryMaxMB		int
	DiskMB			int
}

type TaskRestartPolicy struct {
	Interval		time.Duration
	Attempts 		int
	Delay			int
	Mode 			string
}

type Tasks struct {
	Name 			string
	Driver 			string
	Config 			map[string]interface{}
	Env 			map[string]string
	Resources 		TaskResource
	RestartPolicy	TaskRestartPolicy
}

type Events struct {
	Type 			string
	Time			time.Time
	DisplayMessage 	string
	Message 		string
	FailsTask		bool
	KillReason 		string
	KillTimeout		time.Duration
	KillError 		string
	DownloadError	string
	ValidationError string
	VaultError		string
}

type AllocMemoryStats struct {
	RSS				uint64
	Cache			uint64
	Swap 			uint64
	Usage			uint64
	MaxUsage		uint64
	KernelUsage		uint64
	KernelMaxUsage  uint64
}

type AllocCpuStats struct {
	SystemMode 		float64
	UserMode 		float64
	TotalTicks 		float64
	ThrottledPeriods	uint64
	ThrottledTime 	uint64
	Percent			float64
}

type AllocationStats struct {
	Memory		AllocMemoryStats
	Cpu			AllocCpuStats
}

type JobVersion struct {
	Region 			string
	Namespace 		string
	ID 				string
	Name			string
	Type 			string
	Priority		int
	AllAtOnce		bool
	DataCenters		[]string
	Multiregion		bool
	Stop 			bool
	Status 			string
	Stable 			bool
	Version			uint64
	SubmitTime		time.Time
}

type JobVersionDiff struct {
	Type 			string
	ID				string
	Fields 			[]*api.FieldDiff
	Objects 		[]*api.ObjectDiff
	TaskGroups		[]*api.TaskGroupDiff
}

type TaskGroupDiff struct {
	Type 			string
	Name 			string
	Fields			[]*FieldDiff
	Objects			[]*ObjectDiff
	Tasks			[]*TaskDiff
	Updates			map[string]uint64
}

type TaskDiff struct {
	Type 			string
	Name 			string
	Fields			[]*FieldDiff
	Objects			[]*ObjectDiff
	Annotations		[]string
}

type FieldDiff struct {
	Type 			string
	Name			string
	Old, New		string
	Annotations		[]string
}

type ObjectDiff struct {
	Type 			string
	Name 			string
	Fields 			[]*FieldDiff
	Objects			[]*ObjectDiff
}