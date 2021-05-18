// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	testjobv1 "github.com/kubeflow/common/test_job/apis/test_job/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTestJobs implements TestJobInterface
type FakeTestJobs struct {
	Fake *FakeKubeflowV1
	ns   string
}

var testjobsResource = schema.GroupVersionResource{Group: "kubeflow.org", Version: "v1", Resource: "testjobs"}

var testjobsKind = schema.GroupVersionKind{Group: "kubeflow.org", Version: "v1", Kind: "TestJob"}

// Get takes name of the testJob, and returns the corresponding testJob object, and an error if there is any.
func (c *FakeTestJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *testjobv1.TestJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(testjobsResource, c.ns, name), &testjobv1.TestJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*testjobv1.TestJob), err
}

// List takes label and field selectors, and returns the list of TestJobs that match those selectors.
func (c *FakeTestJobs) List(ctx context.Context, opts v1.ListOptions) (result *testjobv1.TestJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(testjobsResource, testjobsKind, c.ns, opts), &testjobv1.TestJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &testjobv1.TestJobList{ListMeta: obj.(*testjobv1.TestJobList).ListMeta}
	for _, item := range obj.(*testjobv1.TestJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testJobs.
func (c *FakeTestJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(testjobsResource, c.ns, opts))

}

// Create takes the representation of a testJob and creates it.  Returns the server's representation of the testJob, and an error, if there is any.
func (c *FakeTestJobs) Create(ctx context.Context, testJob *testjobv1.TestJob, opts v1.CreateOptions) (result *testjobv1.TestJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(testjobsResource, c.ns, testJob), &testjobv1.TestJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*testjobv1.TestJob), err
}

// Update takes the representation of a testJob and updates it. Returns the server's representation of the testJob, and an error, if there is any.
func (c *FakeTestJobs) Update(ctx context.Context, testJob *testjobv1.TestJob, opts v1.UpdateOptions) (result *testjobv1.TestJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(testjobsResource, c.ns, testJob), &testjobv1.TestJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*testjobv1.TestJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTestJobs) UpdateStatus(ctx context.Context, testJob *testjobv1.TestJob, opts v1.UpdateOptions) (*testjobv1.TestJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(testjobsResource, "status", c.ns, testJob), &testjobv1.TestJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*testjobv1.TestJob), err
}

// Delete takes name of the testJob and deletes it. Returns an error if one occurs.
func (c *FakeTestJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(testjobsResource, c.ns, name), &testjobv1.TestJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTestJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(testjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &testjobv1.TestJobList{})
	return err
}

// Patch applies the patch and returns the patched testJob.
func (c *FakeTestJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *testjobv1.TestJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(testjobsResource, c.ns, name, pt, data, subresources...), &testjobv1.TestJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*testjobv1.TestJob), err
}
