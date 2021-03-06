/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/rancher/gitwatcher/pkg/apis/gitwatcher.cattle.io/v1"
	scheme "github.com/rancher/gitwatcher/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GitCommitsGetter has a method to return a GitCommitInterface.
// A group's client should implement this interface.
type GitCommitsGetter interface {
	GitCommits(namespace string) GitCommitInterface
}

// GitCommitInterface has methods to work with GitCommit resources.
type GitCommitInterface interface {
	Create(*v1.GitCommit) (*v1.GitCommit, error)
	Update(*v1.GitCommit) (*v1.GitCommit, error)
	UpdateStatus(*v1.GitCommit) (*v1.GitCommit, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.GitCommit, error)
	List(opts metav1.ListOptions) (*v1.GitCommitList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.GitCommit, err error)
	GitCommitExpansion
}

// gitCommits implements GitCommitInterface
type gitCommits struct {
	client rest.Interface
	ns     string
}

// newGitCommits returns a GitCommits
func newGitCommits(c *GitwatcherV1Client, namespace string) *gitCommits {
	return &gitCommits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gitCommit, and returns the corresponding gitCommit object, and an error if there is any.
func (c *gitCommits) Get(name string, options metav1.GetOptions) (result *v1.GitCommit, err error) {
	result = &v1.GitCommit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitcommits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GitCommits that match those selectors.
func (c *gitCommits) List(opts metav1.ListOptions) (result *v1.GitCommitList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GitCommitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitcommits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gitCommits.
func (c *gitCommits) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gitcommits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gitCommit and creates it.  Returns the server's representation of the gitCommit, and an error, if there is any.
func (c *gitCommits) Create(gitCommit *v1.GitCommit) (result *v1.GitCommit, err error) {
	result = &v1.GitCommit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gitcommits").
		Body(gitCommit).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gitCommit and updates it. Returns the server's representation of the gitCommit, and an error, if there is any.
func (c *gitCommits) Update(gitCommit *v1.GitCommit) (result *v1.GitCommit, err error) {
	result = &v1.GitCommit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitcommits").
		Name(gitCommit.Name).
		Body(gitCommit).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gitCommits) UpdateStatus(gitCommit *v1.GitCommit) (result *v1.GitCommit, err error) {
	result = &v1.GitCommit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitcommits").
		Name(gitCommit.Name).
		SubResource("status").
		Body(gitCommit).
		Do().
		Into(result)
	return
}

// Delete takes name of the gitCommit and deletes it. Returns an error if one occurs.
func (c *gitCommits) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitcommits").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gitCommits) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitcommits").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gitCommit.
func (c *gitCommits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.GitCommit, err error) {
	result = &v1.GitCommit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gitcommits").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
