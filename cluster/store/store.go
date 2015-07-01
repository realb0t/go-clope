package store

import (
  "fmt"
  "github.com/realb0t/go-clope/cluster"
  "github.com/realb0t/go-clope/transaction"
  "github.com/realb0t/go-clope/cluster/store/driver"
)

type Store struct {
  driver driver.Driver
}

// Create Store instance
func NewStore(drv driver.Driver) *Store {
  return &Store{drv}
}

// Add or move transaction into cluster by clusterId
// and commit changes in store
func (s *Store) MoveTransaction(cId int, t transaction.Transaction) {
  // Если для транзакции был определен кластер
  if t.GetClusterId() != -1 {
    // Удаляем транзакцию из старого кластера
    s.driver.RemoveTransaction(cId, t)
  }
  s.driver.AddTransaction(cId, t)
}

func (s *Store) Driver() driver.Driver {
  return s.driver
}

func (s *Store) Print() {
  clusters, _ := s.driver.Clusters()
  for _, cluster := range(clusters) {
    fmt.Println(cluster)
  }
}

func (s *Store) CreateCluster() (*cluster.Cluster, error) {
  return s.driver.CreateCluster()
}

func (s *Store) RemoveEmpty() error {
  return s.driver.RemoveEmpty()
}

func (s *Store) Iterate(callback func(*cluster.Cluster)) {
  s.driver.Iterate(callback)
}

func (s *Store) Len() int {
  return s.driver.Len()
}
