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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/gateway-api/apisx/v1alpha1"
)

// ListenerSetLister helps list ListenerSets.
// All objects returned here must be treated as read-only.
type ListenerSetLister interface {
	// List lists all ListenerSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ListenerSet, err error)
	// ListenerSets returns an object that can list and get ListenerSets.
	ListenerSets(namespace string) ListenerSetNamespaceLister
	ListenerSetListerExpansion
}

// listenerSetLister implements the ListenerSetLister interface.
type listenerSetLister struct {
	listers.ResourceIndexer[*v1alpha1.ListenerSet]
}

// NewListenerSetLister returns a new ListenerSetLister.
func NewListenerSetLister(indexer cache.Indexer) ListenerSetLister {
	return &listenerSetLister{listers.New[*v1alpha1.ListenerSet](indexer, v1alpha1.Resource("listenerset"))}
}

// ListenerSets returns an object that can list and get ListenerSets.
func (s *listenerSetLister) ListenerSets(namespace string) ListenerSetNamespaceLister {
	return listenerSetNamespaceLister{listers.NewNamespaced[*v1alpha1.ListenerSet](s.ResourceIndexer, namespace)}
}

// ListenerSetNamespaceLister helps list and get ListenerSets.
// All objects returned here must be treated as read-only.
type ListenerSetNamespaceLister interface {
	// List lists all ListenerSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ListenerSet, err error)
	// Get retrieves the ListenerSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ListenerSet, error)
	ListenerSetNamespaceListerExpansion
}

// listenerSetNamespaceLister implements the ListenerSetNamespaceLister
// interface.
type listenerSetNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.ListenerSet]
}
