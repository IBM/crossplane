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

package v1

import (
	"context"
	"time"

	v1 "github.com/crossplane/crossplane/apis/pkg/v1"
	scheme "github.com/crossplane/crossplane/internal/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProvidersGetter has a method to return a ProviderInterface.
// A group's client should implement this interface.
type ProvidersGetter interface {
	Providers(namespace string) ProviderInterface
}

// ProviderInterface has methods to work with Provider resources.
type ProviderInterface interface {
	Create(ctx context.Context, provider *v1.Provider, opts metav1.CreateOptions) (*v1.Provider, error)
	Update(ctx context.Context, provider *v1.Provider, opts metav1.UpdateOptions) (*v1.Provider, error)
	UpdateStatus(ctx context.Context, provider *v1.Provider, opts metav1.UpdateOptions) (*v1.Provider, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Provider, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ProviderList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Provider, err error)
	ProviderExpansion
}

// providers implements ProviderInterface
type providers struct {
	client rest.Interface
	ns     string
}

// newProviders returns a Providers
func newProviders(c *PkgV1Client, namespace string) *providers {
	return &providers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the provider, and returns the corresponding provider object, and an error if there is any.
func (c *providers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Provider, err error) {
	result = &v1.Provider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Providers that match those selectors.
func (c *providers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("providers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested providers.
func (c *providers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("providers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a provider and creates it.  Returns the server's representation of the provider, and an error, if there is any.
func (c *providers) Create(ctx context.Context, provider *v1.Provider, opts metav1.CreateOptions) (result *v1.Provider, err error) {
	result = &v1.Provider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("providers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(provider).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a provider and updates it. Returns the server's representation of the provider, and an error, if there is any.
func (c *providers) Update(ctx context.Context, provider *v1.Provider, opts metav1.UpdateOptions) (result *v1.Provider, err error) {
	result = &v1.Provider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providers").
		Name(provider.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(provider).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *providers) UpdateStatus(ctx context.Context, provider *v1.Provider, opts metav1.UpdateOptions) (result *v1.Provider, err error) {
	result = &v1.Provider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("providers").
		Name(provider.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(provider).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the provider and deletes it. Returns an error if one occurs.
func (c *providers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *providers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("providers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched provider.
func (c *providers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Provider, err error) {
	result = &v1.Provider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("providers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
