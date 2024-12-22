/*
Copyright 2024 The Kubernetes Authors

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

	v1alpha1 "github.com/k8snetworkplumbingwg/whereabouts/pkg/api/whereabouts.cni.cncf.io/v1alpha1"
	scheme "github.com/k8snetworkplumbingwg/whereabouts/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// NodeSlicePoolsGetter has a method to return a NodeSlicePoolInterface.
// A group's client should implement this interface.
type NodeSlicePoolsGetter interface {
	NodeSlicePools(namespace string) NodeSlicePoolInterface
}

// NodeSlicePoolInterface has methods to work with NodeSlicePool resources.
type NodeSlicePoolInterface interface {
	Create(ctx context.Context, nodeSlicePool *v1alpha1.NodeSlicePool, opts v1.CreateOptions) (*v1alpha1.NodeSlicePool, error)
	Update(ctx context.Context, nodeSlicePool *v1alpha1.NodeSlicePool, opts v1.UpdateOptions) (*v1alpha1.NodeSlicePool, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, nodeSlicePool *v1alpha1.NodeSlicePool, opts v1.UpdateOptions) (*v1alpha1.NodeSlicePool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NodeSlicePool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodeSlicePoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeSlicePool, err error)
	NodeSlicePoolExpansion
}

// nodeSlicePools implements NodeSlicePoolInterface
type nodeSlicePools struct {
	*gentype.ClientWithList[*v1alpha1.NodeSlicePool, *v1alpha1.NodeSlicePoolList]
}

// newNodeSlicePools returns a NodeSlicePools
func newNodeSlicePools(c *WhereaboutsV1alpha1Client, namespace string) *nodeSlicePools {
	return &nodeSlicePools{
		gentype.NewClientWithList[*v1alpha1.NodeSlicePool, *v1alpha1.NodeSlicePoolList](
			"nodeslicepools",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.NodeSlicePool { return &v1alpha1.NodeSlicePool{} },
			func() *v1alpha1.NodeSlicePoolList { return &v1alpha1.NodeSlicePoolList{} }),
	}
}
