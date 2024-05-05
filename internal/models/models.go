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

type Metadata struct {
	Host 		string
	Namespace	string
	Region 		string
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
	Version 			*uint64
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

type NodePools struct {
	Name 			string
	Description 	string
	Meta 			map[string]string
	// SchedulerConfig *NodePoolSchedulerConfig
	CreateIndex 	uint64
	ModifyIndex     uint64
}

// type NodePoolSchedulerConfig struct {
// 	SchedulerAlogrithm 				*bool
// 	MemoryOversubscriptionEnabled	*bool
// }

type Nodes struct {
	Address               string
	ID                    string
	Attributes            map[string]string
	Datacenter            string
	Name                  string
	NodeClass             string
	NodePool              string
	Version               string
	Drain                 bool
	SchedulingEligibility string
	Status                string
	StatusDescription     string
	AllocsCount 		  int
	Drivers               map[string]*DriverInfo
	NodeResources         *NodeResources
	ReservedResources     *NodeReservedResources
	LastDrain             *DrainMetadata
	CreateIndex           uint64
	ModifyIndex           uint64
}

type DriverInfo struct {
	Attributes        	map[string]string
	Detected          	bool
	Healthy           	bool
	HealthDescription 	string
	UpdateTime        	time.Time
}

type NodeResources struct {
	Cpu      			NodeCpuResources
	Memory   			NodeMemoryResources
	Disk     			NodeDiskResources
	Networks 			[]*NetworkResource
	Devices  			[]*NodeDeviceResource
	MinDynamicPort 		int
	MaxDynamicPort 		int
}

type NodeCpuResources struct {
	CpuShares          int64
	TotalCpuCores      uint16
	ReservableCpuCores []uint16
}

type NodeMemoryResources struct {
	MemoryMB int64
}

type NodeDiskResources struct {
	DiskMB int64
}

type NetworkResource struct {
	Mode          	string
	Device        	string
	CIDR          	string
	IP            	string
	DNS           	*DNSConfig
	ReservedPorts 	[]Port
	DynamicPorts  	[]Port
	Hostname      	string
	MBits 			*int
}

type DNSConfig struct {
	Servers  		[]string
	Searches 		[]string
	Options  		[]string
}

type Port struct {
	Label       	string 
	Value       	int     
	To          	int
	HostNetwork 	string
}

type NodeDeviceResource struct {
	Vendor 			string
	Type 			string
	Name 			string
	Instances 		[]*NodeDevice
	Attributes 		map[string]*Attribute
}

type NodeDevice struct {
	ID 					string
	Healthy 			bool
	HealthDescription 	string
	Locality 			*NodeDeviceLocality
}

type NodeDeviceLocality struct {
	PciBusID 			string
}

type Attribute struct {
	FloatVal 			*float64
	IntVal 				*int64
	StringVal 			*string
	BoolVal		 		*bool
	Unit 				string
}

type NodeReservedResources struct {
	Cpu      			NodeReservedCpuResources
	Memory   			NodeReservedMemoryResources
	Disk     			NodeReservedDiskResources
	Networks 			NodeReservedNetworkResources
}

type NodeReservedCpuResources struct {
	CpuShares 			uint64
}

type NodeReservedMemoryResources struct {
	MemoryMB 			uint64
}

type NodeReservedDiskResources struct {
	DiskMB 				uint64
}

type NodeReservedNetworkResources struct {
	ReservedHostPorts 	string
}

type DrainMetadata struct {
	StartedAt  			time.Time
	UpdatedAt  			time.Time
	Status     			DrainStatus
	AccessorID 			string
	Meta       			map[string]string
}

type DrainStatus string