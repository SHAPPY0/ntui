package nomad

import (
	// "fmt"
	// "time"
	// "sort"
	"github.com/hashicorp/nomad/api"
	"github.com/shappy0/ntui/internal/models"
	// "github.com/shappy0/ntui/internal/utils"
)

type NodePoolsClient interface {
	List(*api.QueryOptions) ([]*api.NodePool, *api.QueryMeta, error)
	Info(string, *api.QueryOptions) (*api.NodePool, *api.QueryMeta, error)
	ListJobs(string, *api.QueryOptions) ([]*api.JobListStub, *api.QueryMeta, error)
	ListNodes(string, *api.QueryOptions) ([]*api.NodeListStub, *api.QueryMeta, error) 
}

func (n *Nomad) NodePoolList(params *models.NomadParams) ([]models.NodePools, error) {
	result := make([]models.NodePools, 0)
	npList, _, err := n.NodePoolsClient.List(&api.QueryOptions{
		Region: params.Region,
		Namespace: params.Namespace,
	})
	if err != nil {
		n.Logger.Error("Error getting NodePool List: " + err.Error())
		return result, err
	}
	for _, nodePool := range npList {
		var np models.NodePools
		np.Name 		=	nodePool.Name
		np.Description  = 	nodePool.Description
		np.Meta 		=	nodePool.Meta
		np.CreateIndex  =	nodePool.CreateIndex
		np.ModifyIndex  =	nodePool.ModifyIndex
		result = append(result, np)
	}
	return result, nil
}