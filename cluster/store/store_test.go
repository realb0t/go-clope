package store_test

import (
  "testing"
  "github.com/realb0t/go-clope/atom"
  "github.com/realb0t/go-clope/cluster/store"
  "github.com/realb0t/go-clope/transaction/simple"
  "github.com/realb0t/go-clope/cluster/store/driver/memory"
)

func TestNewProcess(t *testing.T) {
  driver := memory.NewMemory()
  _ = store.NewStore(driver)
}

func TestCreateCluster(t *testing.T) {
  driver := memory.NewMemory()
  s := store.NewStore(driver)
  _, _ = s.CreateCluster()
  if s.Len() != 1 {
    t.Error("Clusters count not one")
  }
}

func TestRemoveEmpty(t *testing.T) {
  driver := memory.NewMemory()
  s := store.NewStore(driver)
  trans := simple.Make( "a", "d" )
  _, _ = s.CreateCluster()
  secondCluster, _ := s.CreateCluster()
  s.MoveTransaction(secondCluster.Id, trans)

  if s.Len() != 2 {
    t.Error("Clusters count uncorrect before remove", s.Len())
  }

  _ = s.RemoveEmpty()

  if s.Len() != 1 {
    t.Error("Clusters count uncorrect after remove")
  }
}

func TestMoveTransaction(t *testing.T) {
  driver := memory.NewMemory()
  s := store.NewStore(driver)
  cluster, _ := s.CreateCluster()
  trans1 := simple.Make( "a", "b", "c" )
  trans2 := simple.Make( "a", "d" )
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