// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	machineconfigurationv1 "github.com/openshift/api/machineconfiguration/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// MachineConfigPoolLister helps list MachineConfigPools.
// All objects returned here must be treated as read-only.
type MachineConfigPoolLister interface {
	// List lists all MachineConfigPools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*machineconfigurationv1.MachineConfigPool, err error)
	// Get retrieves the MachineConfigPool from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*machineconfigurationv1.MachineConfigPool, error)
	MachineConfigPoolListerExpansion
}

// machineConfigPoolLister implements the MachineConfigPoolLister interface.
type machineConfigPoolLister struct {
	listers.ResourceIndexer[*machineconfigurationv1.MachineConfigPool]
}

// NewMachineConfigPoolLister returns a new MachineConfigPoolLister.
func NewMachineConfigPoolLister(indexer cache.Indexer) MachineConfigPoolLister {
	return &machineConfigPoolLister{listers.New[*machineconfigurationv1.MachineConfigPool](indexer, machineconfigurationv1.Resource("machineconfigpool"))}
}
