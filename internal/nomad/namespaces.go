package nomad

import (
	// "fmt"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/hashicorp/nomad/api"
)

type NamespaceClient interface {
	List(*api.QueryOptions) ([]*api.Namespace, *api.QueryMeta, error)
}

func (n *Nomad) Namespaces() ([]models.Namespaces, error) {
	result := make([]models.Namespaces, 0)
	data, _, err := n.NamespaceClient.List(nil)
	if err != nil {
		return result, err
	}
	for index, ns := range data {
		var namespace models.Namespaces
		namespace.Id 			= utils.IntToStr(index)
		namespace.Name 			= ns.Name
		namespace.Description 	= ns.Description
		result = append(result, namespace)
	}
	return result, nil
}