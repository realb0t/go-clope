package store_test

import (
  "testing"
  "github.com/realb0t/go-clope/atom"
  "github.com/realb0t/go-clope/cluster/store"
  trn "github.com/realb0t/go-clope/transaction"
)

func TestNewProcess(t *testing.T) {
  _ = store.NewMemoryStore()
}

func TestCreateCluster(t *testing.T) {
  s := store.NewMemoryStore()
  _ = s.CreateCluster()
  if len(s.Clusters) != 1 {
    t.Error("Clusters count not one")
  }
}

func TestRemoveEmpty(t *testing.T) {
  s := store.NewMemoryStore()
  trans := trn.Make( "a", "d" )
  _ = s.CreateCluster()
  secondCluster := s.CreateCluster()
  s.MoveTransaction(secondCluster.Id, trans)

  if len(s.Clusters) != 2 {
    t.Error("Clusters count uncorrect before remove", len(s.Clusters))
  }

  s.RemoveEmpty()

  if len(s.Clusters) != 1 {
    t.Error("Clusters count uncorrect after remove")
  }
}

func TestMoveTransaction(t *testing.T) {
  s := store.NewMemoryStore()
  cluster := s.CreateCluster()
  trans1 := trn.Make( "a", "b", "c" )
  trans2 := trn.Make( "a", "d" )
  s.MoveTransaction(cluster.Id, trans1)
  s.MoveTransaction(cluster.Id, trans2)
  
  if cluster.N != 2 {
    t.Error("Not correct N:", cluster.N)
  }

  if cluster.W != 4 {
    t.Error("Not correct W:", cluster.W)
  }

  if cluster.S != 5 {
    t.Error("Not correct S:", cluster.S)
  }

  cluster.Occ(atom.NewAtom("a"))
}