// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceBindingLister helps list ServiceBindings.
type ServiceBindingLister interface {
	// List lists all ServiceBindings in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error)
	// ServiceBindings returns an object that can list and get ServiceBindings.
	ServiceBindings(namespace string) ServiceBindingNamespaceLister
	ServiceBindingListerExpansion
}

// serviceBindingLister implements the ServiceBindingLister interface.
type serviceBindingLister struct {
	indexer cache.Indexer
}

// NewServiceBindingLister returns a new ServiceBindingLister.
func NewServiceBindingLister(indexer cache.Indexer) ServiceBindingLister {
	return &serviceBindingLister{indexer: indexer}
}

// List lists all ServiceBindings in the indexer.
func (s *serviceBindingLister) List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceBinding))
	})
	return ret, err
}

// ServiceBindings returns an object that can list and get ServiceBindings.
func (s *serviceBindingLister) ServiceBindings(namespace string) ServiceBindingNamespaceLister {
	return serviceBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceBindingNamespaceLister helps list and get ServiceBindings.
type ServiceBindingNamespaceLister interface {
	// List lists all ServiceBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error)
	// Get retrieves the ServiceBinding from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.ServiceBinding, error)
	ServiceBindingNamespaceListerExpansion
}

// serviceBindingNamespaceLister implements the ServiceBindingNamespaceLister
// interface.
type serviceBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceBindings in the indexer for a given namespace.
func (s serviceBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceBinding))
	})
	return ret, err
}

// Get retrieves the ServiceBinding from the indexer for a given namespace and name.
func (s serviceBindingNamespaceLister) Get(name string) (*v1beta1.ServiceBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("servicebinding"), name)
	}
	return obj.(*v1beta1.ServiceBinding), nil
}
