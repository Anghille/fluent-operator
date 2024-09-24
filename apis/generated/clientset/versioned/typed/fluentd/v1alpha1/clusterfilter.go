/*
Copyright 2022.

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

	v1alpha1 "github.com/fluent/fluent-operator/v3/apis/fluentd/v1alpha1"
	scheme "github.com/fluent/fluent-operator/v3/apis/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterFiltersGetter has a method to return a ClusterFilterInterface.
// A group's client should implement this interface.
type ClusterFiltersGetter interface {
	ClusterFilters() ClusterFilterInterface
}

// ClusterFilterInterface has methods to work with ClusterFilter resources.
type ClusterFilterInterface interface {
	Create(ctx context.Context, clusterFilter *v1alpha1.ClusterFilter, opts v1.CreateOptions) (*v1alpha1.ClusterFilter, error)
	Update(ctx context.Context, clusterFilter *v1alpha1.ClusterFilter, opts v1.UpdateOptions) (*v1alpha1.ClusterFilter, error)
	UpdateStatus(ctx context.Context, clusterFilter *v1alpha1.ClusterFilter, opts v1.UpdateOptions) (*v1alpha1.ClusterFilter, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterFilter, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterFilterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterFilter, err error)
	ClusterFilterExpansion
}

// clusterFilters implements ClusterFilterInterface
type clusterFilters struct {
	client rest.Interface
}

// newClusterFilters returns a ClusterFilters
func newClusterFilters(c *FluentdV1alpha1Client) *clusterFilters {
	return &clusterFilters{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterFilter, and returns the corresponding clusterFilter object, and an error if there is any.
func (c *clusterFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterFilter, err error) {
	result = &v1alpha1.ClusterFilter{}
	err = c.client.Get().
		Resource("clusterfilters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterFilters that match those selectors.
func (c *clusterFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterFilterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterFilterList{}
	err = c.client.Get().
		Resource("clusterfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterFilters.
func (c *clusterFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterFilter and creates it.  Returns the server's representation of the clusterFilter, and an error, if there is any.
func (c *clusterFilters) Create(ctx context.Context, clusterFilter *v1alpha1.ClusterFilter, opts v1.CreateOptions) (result *v1alpha1.ClusterFilter, err error) {
	result = &v1alpha1.ClusterFilter{}
	err = c.client.Post().
		Resource("clusterfilters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterFilter).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterFilter and updates it. Returns the server's representation of the clusterFilter, and an error, if there is any.
func (c *clusterFilters) Update(ctx context.Context, clusterFilter *v1alpha1.ClusterFilter, opts v1.UpdateOptions) (result *v1alpha1.ClusterFilter, err error) {
	result = &v1alpha1.ClusterFilter{}
	err = c.client.Put().
		Resource("clusterfilters").
		Name(clusterFilter.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterFilter).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterFilters) UpdateStatus(ctx context.Context, clusterFilter *v1alpha1.ClusterFilter, opts v1.UpdateOptions) (result *v1alpha1.ClusterFilter, err error) {
	result = &v1alpha1.ClusterFilter{}
	err = c.client.Put().
		Resource("clusterfilters").
		Name(clusterFilter.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterFilter).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterFilter and deletes it. Returns an error if one occurs.
func (c *clusterFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterfilters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterfilters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterFilter.
func (c *clusterFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterFilter, err error) {
	result = &v1alpha1.ClusterFilter{}
	err = c.client.Patch(pt).
		Resource("clusterfilters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
