// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	securityv1 "github.com/openshift/api/security/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// SecurityContextConstraintsLister helps list SecurityContextConstraints.
// All objects returned here must be treated as read-only.
type SecurityContextConstraintsLister interface {
	// List lists all SecurityContextConstraints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*securityv1.SecurityContextConstraints, err error)
	// Get retrieves the SecurityContextConstraints from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*securityv1.SecurityContextConstraints, error)
	SecurityContextConstraintsListerExpansion
}

// securityContextConstraintsLister implements the SecurityContextConstraintsLister interface.
type securityContextConstraintsLister struct {
	listers.ResourceIndexer[*securityv1.SecurityContextConstraints]
}

// NewSecurityContextConstraintsLister returns a new SecurityContextConstraintsLister.
func NewSecurityContextConstraintsLister(indexer cache.Indexer) SecurityContextConstraintsLister {
	return &securityContextConstraintsLister{listers.New[*securityv1.SecurityContextConstraints](indexer, securityv1.Resource("securitycontextconstraints"))}
}
