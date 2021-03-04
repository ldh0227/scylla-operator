// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScyllaClusterLister helps list ScyllaClusters.
// All objects returned here must be treated as read-only.
type ScyllaClusterLister interface {
	// List lists all ScyllaClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ScyllaCluster, err error)
	// ScyllaClusters returns an object that can list and get ScyllaClusters.
	ScyllaClusters(namespace string) ScyllaClusterNamespaceLister
	ScyllaClusterListerExpansion
}

// scyllaClusterLister implements the ScyllaClusterLister interface.
type scyllaClusterLister struct {
	indexer cache.Indexer
}

// NewScyllaClusterLister returns a new ScyllaClusterLister.
func NewScyllaClusterLister(indexer cache.Indexer) ScyllaClusterLister {
	return &scyllaClusterLister{indexer: indexer}
}

// List lists all ScyllaClusters in the indexer.
func (s *scyllaClusterLister) List(selector labels.Selector) (ret []*v1.ScyllaCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ScyllaCluster))
	})
	return ret, err
}

// ScyllaClusters returns an object that can list and get ScyllaClusters.
func (s *scyllaClusterLister) ScyllaClusters(namespace string) ScyllaClusterNamespaceLister {
	return scyllaClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScyllaClusterNamespaceLister helps list and get ScyllaClusters.
// All objects returned here must be treated as read-only.
type ScyllaClusterNamespaceLister interface {
	// List lists all ScyllaClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ScyllaCluster, err error)
	// Get retrieves the ScyllaCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ScyllaCluster, error)
	ScyllaClusterNamespaceListerExpansion
}

// scyllaClusterNamespaceLister implements the ScyllaClusterNamespaceLister
// interface.
type scyllaClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScyllaClusters in the indexer for a given namespace.
func (s scyllaClusterNamespaceLister) List(selector labels.Selector) (ret []*v1.ScyllaCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ScyllaCluster))
	})
	return ret, err
}

// Get retrieves the ScyllaCluster from the indexer for a given namespace and name.
func (s scyllaClusterNamespaceLister) Get(name string) (*v1.ScyllaCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("scyllacluster"), name)
	}
	return obj.(*v1.ScyllaCluster), nil
}
