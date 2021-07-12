/*
Copyright 2019 The Crossplane Authors.

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
	"context"

	pkgv1 "github.com/crossplane/crossplane/apis/pkg/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigurations implements ConfigurationInterface
type FakeConfigurations struct {
	Fake *FakePkgV1
	ns   string
}

var configurationsResource = schema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1", Resource: "configurations"}

var configurationsKind = schema.GroupVersionKind{Group: "pkg.crossplane.io", Version: "v1", Kind: "Configuration"}

// Get takes name of the configuration, and returns the corresponding configuration object, and an error if there is any.
func (c *FakeConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *pkgv1.Configuration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configurationsResource, c.ns, name), &pkgv1.Configuration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Configuration), err
}

// List takes label and field selectors, and returns the list of Configurations that match those selectors.
func (c *FakeConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *pkgv1.ConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configurationsResource, configurationsKind, c.ns, opts), &pkgv1.ConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &pkgv1.ConfigurationList{ListMeta: obj.(*pkgv1.ConfigurationList).ListMeta}
	for _, item := range obj.(*pkgv1.ConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configurations.
func (c *FakeConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configurationsResource, c.ns, opts))

}

// Create takes the representation of a configuration and creates it.  Returns the server's representation of the configuration, and an error, if there is any.
func (c *FakeConfigurations) Create(ctx context.Context, configuration *pkgv1.Configuration, opts v1.CreateOptions) (result *pkgv1.Configuration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configurationsResource, c.ns, configuration), &pkgv1.Configuration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Configuration), err
}

// Update takes the representation of a configuration and updates it. Returns the server's representation of the configuration, and an error, if there is any.
func (c *FakeConfigurations) Update(ctx context.Context, configuration *pkgv1.Configuration, opts v1.UpdateOptions) (result *pkgv1.Configuration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configurationsResource, c.ns, configuration), &pkgv1.Configuration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Configuration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigurations) UpdateStatus(ctx context.Context, configuration *pkgv1.Configuration, opts v1.UpdateOptions) (*pkgv1.Configuration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configurationsResource, "status", c.ns, configuration), &pkgv1.Configuration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Configuration), err
}

// Delete takes name of the configuration and deletes it. Returns an error if one occurs.
func (c *FakeConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configurationsResource, c.ns, name), &pkgv1.Configuration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &pkgv1.ConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched configuration.
func (c *FakeConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *pkgv1.Configuration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configurationsResource, c.ns, name, pt, data, subresources...), &pkgv1.Configuration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pkgv1.Configuration), err
}
