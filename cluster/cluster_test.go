package cluster

import (
  "fmt"
  "testing"
  "github.com/realb0t/go-clope/atom"
  trn "github.com/realb0t/go-clope/transaction"
)

func TestNewCluster(t *testing.T) {
  Reset()
  _ = NewCluster(1)
}

func TestAddCluster(t *testing.T) {
  Reset()
  _ = AddCluster()
  if len(Clusters) != 1 {
    t.Error("Clusters count not one")
  }
}

func TestMoveTransaction(t *testing.T) {
  Reset()
  cluster := AddCluster()
  trans1 := trn.BuildTransaction([]string{ "a", "b", "c" })
  trans2 := trn.BuildTransaction([]string{ "a", "d" })
  cluster.MoveTransaction(trans1)
  cluster.MoveTransaction(trans2)
  
  if cluster.N != 2 {
    t.Error("Not correct N:", cluster.N)
  }

  if cluster.W != 4 {
    t.Error("Not correct W:", cluster.W)
  }

  if cluster.S != 5 {
    fmt.Println(cluster.atoms)
    t.Error("Not correct S:", cluster.S)
  }

  cluster.Occ(atom.NewAtom("a"))
}

func TestDeltaAdd(t *testing.T) {
  _ = []*trn.Transaction{ 
    trn.BuildTransaction([]string{ "a", "b" }),
    trn.BuildTransaction([]string{ "a", "b", "c" }),
    trn.BuildTransaction([]string{ "a", "c", "d" }),
    trn.BuildTransaction([]string{ "d", "e" }),
    trn.BuildTransaction([]string{ "d", "e", "f" }),
  }
}
