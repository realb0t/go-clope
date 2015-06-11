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
  // Get Transactions for cluster id
  ClusterTransactions(*cluster.Cluster) []*transaction.Transaction
  Transactions() map[int][]*transaction.Transaction
  Clusters() map[int]*cluster.Cluster
}

type MemoryStore struct {
  clusters map[int]*cluster.Cluster
  transactions map[int][]*transaction.Transaction
  nextId int
}

// Create MemoryStore instance
func NewMemoryStore() *MemoryStore {
  trans := make(map[int][]*transaction.Transaction, 0)
  clusters := make(map[int]*cluster.Cluster, 0)
  return &MemoryStore{clusters, trans, 1}
}

func (s *MemoryStore) Transactions() map[int][]*transaction.Transaction {
  return s.transactions
}

func (s *MemoryStore) Clusters() map[int]*cluster.Cluster {
  return s.clusters
}

func (s *MemoryStore) ClusterTransactions(c *cluster.Cluster) []*transaction.Transaction {
  return s.transactions[c.Id]
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
    s.RemoveTransaction(cId, t)
  }
  s.AddTransaction(cId, t)
}

// Remove transaction from cluster
func (s *MemoryStore) RemoveTransaction(cId int, t *transaction.Transaction) {
  ei := -1

  // Опеределяем индекс данной транзакции
  for i, trans := range(s.transactions[cId]) {
    if (t == trans) { ei = i }
  }

  if ei != -1 {
    // Извлекаем ее из массива транзакций
    copy(s.transactions[cId][ei:], s.transactions[cId][ei+1:])
    s.transactions[cId][len(s.transactions[cId])-1] = nil
    s.transactions[cId] = s.transactions[cId][:len(s.transactions[cId])-1]

    s.clusters[cId].RefreshAfterRemove(t)
  }
}

// Add transaction into cluster
func (s *MemoryStore) AddTransaction(cId int, t *transaction.Transaction) {
  s.transactions[cId] = append(s.transactions[cId], t)
  s.clusters[cId].RefreshAfterAdd(t)
  t.ClusterId = cId
}

func (s *MemoryStore) RemoveEmpty() {
  for id, cluster := range(s.clusters) {
    if len(s.transactions[cluster.Id]) == 0 {
      delete(s.clusters, id)
    }
  }
}

func (s *MemoryStore) Print() {
  for _, cluster := range(s.clusters) {
    fmt.Println(cluster)
  }
}
