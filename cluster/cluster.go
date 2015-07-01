package cluster

import (
  "fmt"
  "math"
  "github.com/realb0t/go-clope/atom"
  trn "github.com/realb0t/go-clope/transaction"
)

// Cluster struct
type Cluster struct {
  Id int // ID
  // Cluster parameters
  N int // Transaction quantity
  W int // Quantity of uniq atoms (cluster width) 
  S int // Area of cluster
  // Map of atoms
  // Key - Atom object
  // Value - Atoms quantity
  atoms map[*atom.Atom]int
}

// Create new enmpty cluster
func NewCluster(id int) *Cluster {
  return &Cluster{id, 0, 0, 0, make(map[*atom.Atom]int, 0)}
}

// Inspect cluster
func (c *Cluster) String() string {
  s := fmt.Sprintf("ID:%d;N:%d;W:%d;S:%d;%v", c.Id, c.N, c.W, c.S, c.atoms)
  return s
}

// Atom quantity
func (c *Cluster) Occ(atom *atom.Atom) int {
  return c.atoms[atom]
}

// Refresh cluster struct after add transaction
func (c *Cluster) RefreshAfterAdd(t trn.Transaction) {
  atoms := t.Atoms()
  for _, atom := range(atoms) {
    if count, ok := c.atoms[atom] ; ok {
      c.atoms[atom] = count + 1
    } else {
      c.atoms[atom] = 1
    }
  }
  c.refresh(c.N + 1)
}

// Refresh cluster struct after remove transaction
func (c *Cluster) RefreshAfterRemove(t trn.Transaction) {
  atoms := t.Atoms()
  for _, atom := range(atoms) {
    if count, ok := c.atoms[atom] ; ok {
      c.atoms[atom] = count - 1
      if c.atoms[atom] == 0 {
        delete(c.atoms, atom)
      }
    }
  }
  c.refresh(c.N - 1)
}

// Refresh cluster parameters
func (c *Cluster) refresh(transCount int) {
  c.N = transCount
  c.W = len(c.atoms)
  c.S = 0

  for _, count := range(c.atoms) {
    c.S = c.S + count
  }
}

// DeltaAdd calculation
func (c *Cluster) DeltaAdd(t trn.Transaction, r float64) float64 {
  atoms           := t.Atoms()
  transAtomsCount := len(atoms)
  S_new           := float64(c.S + transAtomsCount)
  W_new           := float64(c.W)
  toCounter       := transAtomsCount - 1

  for i := 0; i < toCounter; i++ {
    if _, ok := c.atoms[atoms[i]]; !ok { W_new++; }
  }

  if c.N == 0 {
    return S_new / math.Pow(W_new, r)
  }

  profitCur := float64(c.S) * float64(c.N) / math.Pow(float64(c.W), r)
  profitNew := S_new * float64(c.N + 1) / math.Pow(W_new, r)
  return profitNew - profitCur
}
