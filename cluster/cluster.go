package cluster

import (
  "math"
  "github.com/realb0t/go-clope/atom"
  "github.com/realb0t/go-clope/transaction"
)

type Cluster struct {
  transactions []*transaction.Transaction
  N int
  W int
  S float64
  r float
}

func (c *Cluster) Occ(_ *atom.Atom) float64 {
  return 0.0
}

func (c *Cluster) DeltaAdd(t *Transaction) float64 {
  tItemsCount := count(t.Items)
  S_new := c.S + tItemsCount
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
