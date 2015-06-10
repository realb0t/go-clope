package store

import (
  "fmt"
  "github.com/realb0t/go-clope/cluster"
  "github.com/realb0t/go-clope/transaction"
)

type ClusterStore interface {
  // Create new cluster in store and commit in store
  CreateCluster() *cluster.Cluster
  // Add or move transaction into cluster by clusterId
  // and commit changes in store
  MoveTransaction(clusterId int, trans *transaction.Transaction)
  // Remove all empty clusters from store
  RemoveEmpty()
  // Iterate all clusters
  // @todo Rename to IterateClusters
  Iterate(callback func(*cluster.Cluster))
  // Length all clusters
  Len() int
  // Output all store
  Print()
  // Get cluster by id
  Cluster(id int) *cluster.Cluster
}

type MemoryStore struct {
  clusters map[int]*cluster.Cluster
  nextId int
}

// Create MemoryStore instance
func NewMemoryStore() *MemoryStore {
  return &MemoryStore{make(map[int]*cluster.Cluster, 0), 1}
}

func (s *MemoryStore) Iterate(callback func(*cluster.Cluster)) {
  for _, cluster := range(s.clusters) {
    callback(cluster)
  }
}

func (s *MemoryStore) Cluster(id int) *cluster.Cluster {
  return s.clusters[id]
}

func (s *MemoryStore) Len() int {
  return len(s.clusters)
}

// Create new cluster in store and commit in store
func (s *MemoryStore) CreateCluster() *cluster.Cluster {
  curId := s.nextId
  s.nextId++
  s.clusters[curId] = cluster.NewCluster(curId)
  return s.clusters[curId]
}

  // Add or move transaction into cluster by clusterId
  // and commit changes in store
func (s *MemoryStore) MoveTransaction(cId int, t *transaction.Transaction) {
  // Если для транзакции был определен кластер
  if t.ClusterId != -1 {
    // Удаляем транзакцию из старого кластера
    s.clusters[t.ClusterId].RemoveTransaction(t)
  }
  s.clusters[cId].AddTransaction(t)
}

func (s *MemoryStore) RemoveEmpty() {
  for id, cluster := range(s.clusters) {
    if cluster.IsEmpty() {
      delete(s.clusters, id)
    }
  }
}

func (s *MemoryStore) Print() {
  for _, cluster := range(s.clusters) {
    fmt.Println(cluster)
  }
}
