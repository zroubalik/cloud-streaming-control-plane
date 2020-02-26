/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/streaming/pkg/apis/streaming/v1alpha1"
)

// FakeInbounds implements InboundInterface
type FakeInbounds struct {
	Fake *FakeStreamingV1alpha1
	ns   string
}

var inboundsResource = schema.GroupVersionResource{Group: "streaming.knative.dev", Version: "v1alpha1", Resource: "inbounds"}

var inboundsKind = schema.GroupVersionKind{Group: "streaming.knative.dev", Version: "v1alpha1", Kind: "Inbound"}

// Get takes name of the inbound, and returns the corresponding inbound object, and an error if there is any.
func (c *FakeInbounds) Get(name string, options v1.GetOptions) (result *v1alpha1.Inbound, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(inboundsResource, c.ns, name), &v1alpha1.Inbound{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Inbound), err
}

// List takes label and field selectors, and returns the list of Inbounds that match those selectors.
func (c *FakeInbounds) List(opts v1.ListOptions) (result *v1alpha1.InboundList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(inboundsResource, inboundsKind, c.ns, opts), &v1alpha1.InboundList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InboundList{ListMeta: obj.(*v1alpha1.InboundList).ListMeta}
	for _, item := range obj.(*v1alpha1.InboundList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inbounds.
func (c *FakeInbounds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(inboundsResource, c.ns, opts))

}

// Create takes the representation of a inbound and creates it.  Returns the server's representation of the inbound, and an error, if there is any.
func (c *FakeInbounds) Create(inbound *v1alpha1.Inbound) (result *v1alpha1.Inbound, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(inboundsResource, c.ns, inbound), &v1alpha1.Inbound{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Inbound), err
}

// Update takes the representation of a inbound and updates it. Returns the server's representation of the inbound, and an error, if there is any.
func (c *FakeInbounds) Update(inbound *v1alpha1.Inbound) (result *v1alpha1.Inbound, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(inboundsResource, c.ns, inbound), &v1alpha1.Inbound{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Inbound), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInbounds) UpdateStatus(inbound *v1alpha1.Inbound) (*v1alpha1.Inbound, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(inboundsResource, "status", c.ns, inbound), &v1alpha1.Inbound{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Inbound), err
}

// Delete takes name of the inbound and deletes it. Returns an error if one occurs.
func (c *FakeInbounds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(inboundsResource, c.ns, name), &v1alpha1.Inbound{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInbounds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(inboundsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InboundList{})
	return err
}

// Patch applies the patch and returns the patched inbound.
func (c *FakeInbounds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Inbound, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(inboundsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Inbound{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Inbound), err
}