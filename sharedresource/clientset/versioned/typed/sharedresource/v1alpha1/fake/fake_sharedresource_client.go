// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/client-go/sharedresource/clientset/versioned/typed/sharedresource/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSharedresourceV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSharedresourceV1alpha1) SharedConfigMaps() v1alpha1.SharedConfigMapInterface {
	return newFakeSharedConfigMaps(c)
}

func (c *FakeSharedresourceV1alpha1) SharedSecrets() v1alpha1.SharedSecretInterface {
	return newFakeSharedSecrets(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSharedresourceV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
