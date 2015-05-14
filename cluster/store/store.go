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
  MoveTransaction(clusterId int, trans *transaction.Transaction) *cluster.Cluster
  // Remove all empty clusters from store
  RemoveEmpty()
}

type MemoryStore struct {
  Clusters map[int]*cluster.Cluster
  nextId int
}

func NewMemoryStore() *MemoryStore {
  return &MemoryStore{make(map[int]*cluster.Cluster, 0), 1}
}

func (s *MemoryStore) CreateCluster() *cluster.Cluster {
  curId := s.nextId
  s.nextId++
  s.Clusters[curId] = cluster.NewCluster(curId)
  return s.Clusters[curId]
}

func (s *MemoryStore) MoveTransaction(cId int, t *transaction.Transaction) {
  // Если для транзакции был определен кластер
  if t.ClusterId != -1 {
    // Удаляем транзакцию из старого кластера
    s.Clusters[t.ClusterId].RemoveTransaction(t)
  }
  s.Clusters[cId].AddTransaction(t)
}

func (s *MemoryStore) RemoveEmpty() {
  for id, cluster := range(s.Clusters) {
    if cluster.IsEmpty() {
      delete(s.Clusters, id)
    }
  }
}

func (s *MemoryStore) Print() {
  for _, cluster := range(s.Clusters) {
    fmt.Println(cluster)
  }
}
