package driver

import (
  "github.com/realb0t/go-clope/cluster"
  "github.com/realb0t/go-clope/transaction"
)

type Driver interface {
  // Create new cluster in store and commit in store
  CreateCluster() (*cluster.Cluster, error)
  RemoveTransaction(cId int, t transaction.Transaction)
  AddTransaction(cId int, t transaction.Transaction)
  // Remove all empty clusters from store
  RemoveEmpty() error
  Iterate(callback func(*cluster.Cluster))
  // Length all clusters
  Len() int
  // Get cluster by id
  Cluster(id int) (*cluster.Cluster, error)
  // Get Transactions for Cluster
  ClusterTransactions(*cluster.Cluster) ([]transaction.Transaction, error)
  // Get All Transactions for All clusters as map
  Transactions() (map[int][]transaction.Transaction, error)
  // Get All Clusters
  Clusters() (map[int]*cluster.Cluster, error)
}