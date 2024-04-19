package nomad

import (
	// "fmt"
	// "time"
	// "sort"
	"github.com/hashicorp/nomad/api"
	"github.com/shappy0/ntui/internal/models"
	// "github.com/shappy0/ntui/internal/utils"
)

type NodesClient interface {
	List(*api.QueryOptions) ([]*api.NodeListStub, *api.QueryMeta, error)
	Allocations(string, *api.QueryOptions) ([]*api.Allocation, *api.QueryMeta, error)
}

func (n *Nomad) NodeList(params *models.NomadParams) ([]models.Nodes, error) {
	result := make([]models.Nodes, 0)
	nList, _, err := n.NodesClient.List(&api.QueryOptions{
		Region: params.Region,
		Namespace: params.Namespace,
	})
	if err != nil {
		n.Logger.Errorf("Error getting Node List: %s", err.Error())
		return result, err
	}
	for _, node := range nList {
		n := models.Nodes{
			Address: 				node.Address,
			ID:						node.ID,
			Attributes: 			map[string]string{},
			Datacenter: 			node.Datacenter,
			Name:					node.Name,
			NodeClass:				node.NodeClass,
			NodePool:				node.NodePool,
			Version:				node.Version,
			Drain:					node.Drain,
			SchedulingEligibility: 	node.SchedulingEligibility,
			Status:					node.Status,
			AllocsCount: 			n.GetAllocCounts(node.ID, params),
			StatusDescription: 		node.StatusDescription,
			// Drivers: 			node.Drivers,
			// LastDrain:  			node.LastDrain,
			CreateIndex: 			node.CreateIndex,
			ModifyIndex: 			node.ModifyIndex,
		}
		result = append(result, n) 
	}
	return result, nil
}

func (n *Nomad) GetAllocCounts(nodeId string, params *models.NomadParams) int {
	if nodeId == "" {
		return 0
	}
	allocs, _, err := n.NodesClient.Allocations(nodeId, &api.QueryOptions{
		Region: params.Region,
		Namespace: params.Namespace,
	})
	if err != nil {
		n.Logger.Errorf("Error getting Alloc Counts: %s", err.Error())
		return 0
	}
	return len(allocs)
}