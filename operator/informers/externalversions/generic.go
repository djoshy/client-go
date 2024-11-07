// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/openshift/api/operator/v1"
	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=operator.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("authentications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().Authentications().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("csisnapshotcontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().CSISnapshotControllers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("cloudcredentials"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().CloudCredentials().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("clustercsidrivers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().ClusterCSIDrivers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("configs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().Configs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("consoles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().Consoles().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("dnses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().DNSes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("etcds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().Etcds().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ingresscontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().IngressControllers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("insightsoperators"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().InsightsOperators().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("kubeapiservers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().KubeAPIServers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("kubecontrollermanagers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().KubeControllerManagers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("kubeschedulers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().KubeSchedulers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("kubestorageversionmigrators"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().KubeStorageVersionMigrators().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("machineconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().MachineConfigurations().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("networks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().Networks().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("olms"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().OLMs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("openshiftapiservers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().OpenShiftAPIServers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("openshiftcontrollermanagers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().OpenShiftControllerManagers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("servicecas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().ServiceCAs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("servicecatalogapiservers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().ServiceCatalogAPIServers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("servicecatalogcontrollermanagers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().ServiceCatalogControllerManagers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("storages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1().Storages().Informer()}, nil

		// Group=operator.openshift.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("etcdbackups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().EtcdBackups().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("imagecontentsourcepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().ImageContentSourcePolicies().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("olms"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().OLMs().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
