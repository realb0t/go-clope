package memory

import (
  "github.com/realb0t/go-clope/cluster"
  "github.com/realb0t/go-clope/transaction"
)

type Memory struct {
  clusters map[int]*cluster.Cluster
  transactions map[int][]*transaction.Transaction
  nextId int
}

// Create Memory instance
func NewMemory() *Memory {
  trans := make(map[int][]*transaction.Transaction, 0)
  clusters := make(map[int]*cluster.Cluster, 0)
  return &Memory{clusters, trans, 1}
}

func (s *Memory) Transactions() (map[int][]*transaction.Transaction, error) {
  return s.transactions, nil
}

func (s *Memory) Clusters() (map[int]*cluster.Cluster, error) {
  return s.clusters, nil
}

func (s *Memory) ClusterTransactions(c *cluster.Cluster) ([]*transaction.Transaction, error) {
  return s.transactions[c.Id], nil
}

func (s *Memory) Iterate(callback func(*cluster.Cluster)) {
  for _, cluster := range(s.clusters) {
    callback(cluster)
  }
}

func (s *Memory) Cluster(id int) (*cluster.Cluster, error) {
  return s.clusters[id], nil
}

func (s *Memory) Len() int {
  return len(s.clusters)
}

// Create new cluster in store and commit in store
func (s *Memory) CreateCluster() (*cluster.Cluster, error) {
  curId := s.nextId
  s.nextId++
  s.clusters[curId] = cluster.NewCluster(curId)
  return s.clusters[curId], nil
}

// Remove transaction from cluster
func (s *Memory) RemoveTransaction(cId int, t *transaction.Transaction) {
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
func (s *Memory) AddTransaction(cId int, t *transaction.Transaction) {
  s.transactions[cId] = append(s.transactions[cId], t)
  s.clusters[cId].RefreshAfterAdd(t)
  t.ClusterId = cId
}

func (s *Memory) RemoveEmpty() error {
  for id, cluster := range(s.clusters) {
    if len(s.transactions[cluster.Id]) == 0 {
      delete(s.clusters, id)
    }
  }

  return nil
}

