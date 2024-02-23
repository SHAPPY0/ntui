package nomad

import (
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

type RegionClient interface {
	List() ([]string, error)
}

func (n *Nomad) Regions() ([]models.Regions, error) {
	Result := make([]models.Regions, 0)
	Data, Err := n.RegionClient.List()
	if Err != nil {
		return Result, Err
	}
	for Index, Name := range Data {
		var Region models.Regions
		Region.Id 		=	utils.IntToStr(Index)
		Region.Name 	=	Name
		Result = append(Result, Region)
	}
	return Result, nil
}