package nomad

import (
	// "fmt"
	"github.com/hashicorp/nomad/api"
	"github.com/shappy0/ntui/internal/utils"
)

type Nomad struct {
	Client 				*api.Client
	RegionClient 		RegionClient
	NamespaceClient		NamespaceClient
	JobClient 			JobClient
	AllocationClient 	AllocationClient
	AllocationFSClient 	AllocationFSClient
	Namespace 			string
	Logger 				*utils.Logger	
}

func New(logger *utils.Logger) (*Nomad, error) {
	n := &Nomad{
		Logger:		logger,
	}
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic("Error connecting nomad client: " + err.Error())
	}
	n.Client = client
	n.RegionClient = n.Client.Regions()
	n.NamespaceClient = n.Client.Namespaces()
	n.JobClient = n.Client.Jobs()
	n.AllocationClient = n.Client.Allocations()
	n.AllocationFSClient = n.Client.AllocFS()
	return n, nil
}