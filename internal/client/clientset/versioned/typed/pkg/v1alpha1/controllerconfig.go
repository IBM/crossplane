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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/crossplane/crossplane/apis/pkg/v1alpha1"
	scheme "github.com/crossplane/crossplane/internal/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ControllerConfigsGetter has a method to return a ControllerConfigInterface.
// A group's client should implement this interface.
type ControllerConfigsGetter interface {
	ControllerConfigs(namespace string) ControllerConfigInterface
}

// ControllerConfigInterface has methods to work with ControllerConfig resources.
type ControllerConfigInterface interface {
	Create(ctx context.Context, controllerConfig *v1alpha1.ControllerConfig, opts v1.CreateOptions) (*v1alpha1.ControllerConfig, error)
	Update(ctx context.Context, controllerConfig *v1alpha1.ControllerConfig, opts v1.UpdateOptions) (*v1alpha1.ControllerConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ControllerConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ControllerConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ControllerConfig, err error)
	ControllerConfigExpansion
}

// controllerConfigs implements ControllerConfigInterface
type controllerConfigs struct {
	client rest.Interface
	ns     string
}

// newControllerConfigs returns a ControllerConfigs
func newControllerConfigs(c *PkgV1alpha1Client, namespace string) *controllerConfigs {
	return &controllerConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the controllerConfig, and returns the corresponding controllerConfig object, and an error if there is any.
func (c *controllerConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ControllerConfig, err error) {
	result = &v1alpha1.ControllerConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("controllerconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ControllerConfigs that match those selectors.
func (c *controllerConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ControllerConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ControllerConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("controllerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested controllerConfigs.
func (c *controllerConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("controllerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a controllerConfig and creates it.  Returns the server's representation of the controllerConfig, and an error, if there is any.
func (c *controllerConfigs) Create(ctx context.Context, controllerConfig *v1alpha1.ControllerConfig, opts v1.CreateOptions) (result *v1alpha1.ControllerConfig, err error) {
	result = &v1alpha1.ControllerConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("controllerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(controllerConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a controllerConfig and updates it. Returns the server's representation of the controllerConfig, and an error, if there is any.
func (c *controllerConfigs) Update(ctx context.Context, controllerConfig *v1alpha1.ControllerConfig, opts v1.UpdateOptions) (result *v1alpha1.ControllerConfig, err error) {
	result = &v1alpha1.ControllerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("controllerconfigs").
		Name(controllerConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(controllerConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the controllerConfig and deletes it. Returns an error if one occurs.
func (c *controllerConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("controllerconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *controllerConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("controllerconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched controllerConfig.
func (c *controllerConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ControllerConfig, err error) {
	result = &v1alpha1.ControllerConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("controllerconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
