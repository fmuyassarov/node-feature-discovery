/*
Copyright 2025 The Kubernetes Authors.

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
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	nfdv1alpha1 "sigs.k8s.io/node-feature-discovery/api/nfd/v1alpha1"
)

// NodeFeatureRuleLister helps list NodeFeatureRules.
// All objects returned here must be treated as read-only.
type NodeFeatureRuleLister interface {
	// List lists all NodeFeatureRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*nfdv1alpha1.NodeFeatureRule, err error)
	// Get retrieves the NodeFeatureRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*nfdv1alpha1.NodeFeatureRule, error)
	NodeFeatureRuleListerExpansion
}

// nodeFeatureRuleLister implements the NodeFeatureRuleLister interface.
type nodeFeatureRuleLister struct {
	listers.ResourceIndexer[*nfdv1alpha1.NodeFeatureRule]
}

// NewNodeFeatureRuleLister returns a new NodeFeatureRuleLister.
func NewNodeFeatureRuleLister(indexer cache.Indexer) NodeFeatureRuleLister {
	return &nodeFeatureRuleLister{listers.New[*nfdv1alpha1.NodeFeatureRule](indexer, nfdv1alpha1.Resource("nodefeaturerule"))}
}
