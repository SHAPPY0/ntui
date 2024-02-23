package nomad

import (
	"fmt"
	"github.com/hashicorp/nomad/api"
)

type Nomad struct {
	Client 				*api.Client
	RegionClient 		RegionClient
	NamespaceClient		NamespaceClient
	JobClient 			JobClient
	AllocationClient 	AllocationClient
	AllocationFSClient 	AllocationFSClient
	Namespace 			string
}

func New() (*Nomad, error) {
	Nomad := &Nomad{}
	Client, Err := api.NewClient(api.DefaultConfig())
	if Err != nil {
		fmt.Println("Error while creating nomad client ", Err)
		return Nomad, Err
	}
	Nomad.Client = Client
	Nomad.RegionClient = Nomad.Client.Regions()
	Nomad.NamespaceClient = Nomad.Client.Namespaces()
	Nomad.JobClient = Nomad.Client.Jobs()
	Nomad.AllocationClient = Nomad.Client.Allocations()
	Nomad.AllocationFSClient = Nomad.Client.AllocFS()
	return Nomad, nil
}