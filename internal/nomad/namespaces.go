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
	Result := make([]models.Namespaces, 0)
	Data, _, Err := n.NamespaceClient.List(nil)
	if Err != nil {
		return Result, Err
	}
	for Index, Ns := range Data {
		var Namespace models.Namespaces
		Namespace.Id 			= utils.IntToStr(Index)
		Namespace.Name 			= Ns.Name
		Namespace.Description 	= Ns.Description
		Result = append(Result, Namespace)
	}
	return Result, nil
}