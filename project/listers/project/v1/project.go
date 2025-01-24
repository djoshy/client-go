// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	projectv1 "github.com/openshift/api/project/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ProjectLister helps list Projects.
// All objects returned here must be treated as read-only.
type ProjectLister interface {
	// List lists all Projects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*projectv1.Project, err error)
	// Get retrieves the Project from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*projectv1.Project, error)
	ProjectListerExpansion
}

// projectLister implements the ProjectLister interface.
type projectLister struct {
	listers.ResourceIndexer[*projectv1.Project]
}

// NewProjectLister returns a new ProjectLister.
func NewProjectLister(indexer cache.Indexer) ProjectLister {
	return &projectLister{listers.New[*projectv1.Project](indexer, projectv1.Resource("project"))}
}
