// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/network/v1"
	networkv1 "github.com/openshift/client-go/network/applyconfigurations/network/v1"
	typednetworkv1 "github.com/openshift/client-go/network/clientset/versioned/typed/network/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeHostSubnets implements HostSubnetInterface
type fakeHostSubnets struct {
	*gentype.FakeClientWithListAndApply[*v1.HostSubnet, *v1.HostSubnetList, *networkv1.HostSubnetApplyConfiguration]
	Fake *FakeNetworkV1
}

func newFakeHostSubnets(fake *FakeNetworkV1) typednetworkv1.HostSubnetInterface {
	return &fakeHostSubnets{
		gentype.NewFakeClientWithListAndApply[*v1.HostSubnet, *v1.HostSubnetList, *networkv1.HostSubnetApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("hostsubnets"),
			v1.SchemeGroupVersion.WithKind("HostSubnet"),
			func() *v1.HostSubnet { return &v1.HostSubnet{} },
			func() *v1.HostSubnetList { return &v1.HostSubnetList{} },
			func(dst, src *v1.HostSubnetList) { dst.ListMeta = src.ListMeta },
			func(list *v1.HostSubnetList) []*v1.HostSubnet { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.HostSubnetList, items []*v1.HostSubnet) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
