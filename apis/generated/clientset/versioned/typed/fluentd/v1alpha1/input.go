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

// InputsGetter has a method to return a InputInterface.
// A group's client should implement this interface.
type InputsGetter interface {
	Inputs(namespace string) InputInterface
}

// InputInterface has methods to work with Input resources.
type InputInterface interface {
	Create(ctx context.Context, input *v1alpha1.Input, opts v1.CreateOptions) (*v1alpha1.Input, error)
	Update(ctx context.Context, input *v1alpha1.Input, opts v1.UpdateOptions) (*v1alpha1.Input, error)
	UpdateStatus(ctx context.Context, input *v1alpha1.Input, opts v1.UpdateOptions) (*v1alpha1.Input, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Input, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.InputList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Input, err error)
	InputExpansion
}

// inputs implements InputInterface
type inputs struct {
	client rest.Interface
	ns     string
}

// newInputs returns a Inputs
func newInputs(c *FluentdV1alpha1Client, namespace string) *inputs {
	return &inputs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the input, and returns the corresponding input object, and an error if there is any.
func (c *inputs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Input, err error) {
	result = &v1alpha1.Input{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inputs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Inputs that match those selectors.
func (c *inputs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InputList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.InputList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inputs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested inputs.
func (c *inputs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("inputs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a input and creates it.  Returns the server's representation of the input, and an error, if there is any.
func (c *inputs) Create(ctx context.Context, input *v1alpha1.Input, opts v1.CreateOptions) (result *v1alpha1.Input, err error) {
	result = &v1alpha1.Input{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("inputs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(input).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a input and updates it. Returns the server's representation of the input, and an error, if there is any.
func (c *inputs) Update(ctx context.Context, input *v1alpha1.Input, opts v1.UpdateOptions) (result *v1alpha1.Input, err error) {
	result = &v1alpha1.Input{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inputs").
		Name(input.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(input).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *inputs) UpdateStatus(ctx context.Context, input *v1alpha1.Input, opts v1.UpdateOptions) (result *v1alpha1.Input, err error) {
	result = &v1alpha1.Input{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inputs").
		Name(input.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(input).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the input and deletes it. Returns an error if one occurs.
func (c *inputs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inputs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *inputs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inputs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched input.
func (c *inputs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Input, err error) {
	result = &v1alpha1.Input{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("inputs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
