package cluster

import (
  "math"
  "github.com/realb0t/go-clope/atom"
  trans "github.com/realb0t/go-clope/transaction"
)

type Cluster struct {
  transactions []*trans.Transaction
  N int
  W int
  S float64
}

func NewCluster() {
  return &Cluster{[], 0, 0, 0.0}
}

func (c *Cluster) Occ(_ *atom.Atom) float64 {
  return 0.0
}

func (c *Cluster) AddTransaction(t *trans.Transaction) {
  c.transactions = append(c.transactions, t)
  t.Cluster.RemTransaction(t)
  t.Cluster = c

  c.N = len(c.transactions)
  c.W = 0
}

func (c *Cluster) RemTransaction(t *trans.Transaction) {
  ei := -1

  for i, trans := range(c.transactions) {
    if (t == trans) {  ei = i }
  }

  copy(c.transactions[ei:], c.transactions[ei+1:])
  c.transactions[len(c.transactions)-1] = nil
  c.transactions = c.transactions[:len(c.transactions)-1]
}

func (c *Cluster) DeltaAdd(t *Transaction) float64 {
  tItemsCount := count(t.Items)
  S_new := c.S + float64(tItemsCount)
  W_new := c.W
  toCounter := tItemsCount - 1
  for i := 0; i < toCounter; i++ {
    if c.Occ(t.Items[i]) == nil { W_new++; }
  }

  if c.N == 0 {
    return S_new / Math.pow(W_new, r)
  }

  profitCur := c.S * c.N / Math.pow(c.W, r)
  profitNew := S_new * (c.N + 1) / Math.pow(W_new, r)
  return profitNew - profitCur
}
