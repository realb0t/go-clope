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

func TestRemoveEmpty(t *testing.T) {
  Reset()
  trans := trn.Make( "a", "d" )
  _ = AddCluster()
  secondCluster := AddCluster()
  secondCluster.MoveTransaction(trans)

  if len(Clusters) != 2 {
    t.Error("Clusters count uncorrect before remove", len(Clusters))
  }

  RemoveEmpty()

  if len(Clusters) != 1 {
    t.Error("Clusters count uncorrect after remove")
  }
}

func TestMoveTransaction(t *testing.T) {
  Reset()
  cluster := AddCluster()
  trans1 := trn.Make( "a", "b", "c" )
  trans2 := trn.Make( "a", "d" )
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

func TestDeltaAddEvaluative(t *testing.T) {
  r := 2.6

  trans := []*trn.Transaction{
    trn.Make( "a", "b" ),
    trn.Make( "a", "b", "c" ),
    trn.Make( "a", "c", "d" ),
    trn.Make( "d", "e" ),
    trn.Make( "d", "e", "f" ),
  }

  clusters := []*Cluster{ AddCluster(),
    AddCluster(), AddCluster(), AddCluster() }

  clusters[0].MoveTransaction(trans[0])
  clusters[0].MoveTransaction(trans[1])
  clusters[0].MoveTransaction(trans[2])
  clusters[1].MoveTransaction(trans[3])
  clusters[1].MoveTransaction(trans[4])
  clusters[2].MoveTransaction(trans[0])
  clusters[2].MoveTransaction(trans[1])
  clusters[3].MoveTransaction(trans[2])
  clusters[3].MoveTransaction(trans[3])
  clusters[3].MoveTransaction(trans[4])

  if clusters[0].DeltaAdd(trans[0], r) <= clusters[3].DeltaAdd(trans[0], r) {
    t.Error("Bad DeltaAdd for first compare")
  }

  if clusters[3].DeltaAdd(trans[3], r) <= clusters[2].DeltaAdd(trans[3], r) {
    t.Error("Bad DeltaAdd for second compare")
  }

  if clusters[2].DeltaAdd(trans[2], r) <= clusters[3].DeltaAdd(trans[2], r) {
    t.Error("Bad DeltaAdd for third compare")
  }
}
