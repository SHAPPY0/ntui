package nomad

import (
	"github.com/hashicorp/nomad/api"
)

type AllocationFSClient interface {
	Logs(
		alloc *api.Allocation,
		follow bool,
		task string,
		logType string,
		origin 	string,
		offset int64,
		cancel <-chan struct{},
		query *api.QueryOptions) (<-chan *api.StreamFrame, <-chan error)
}

func (n *Nomad) Logs(allocID, taskName, logType, origin string, follow bool, offset int64, cancel <-chan struct{}) (<-chan *api.StreamFrame, <-chan error) {
	return n.AllocationFSClient.Logs(&api.Allocation{ID: allocID},
					follow,
					taskName,
					logType,
					origin,
					offset,
					cancel,
					&api.QueryOptions{},
	)
}