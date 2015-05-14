package cluster

import (
  "testing"
  trn "github.com/realb0t/go-clope/transaction"
)

func TestNewCluster(t *testing.T) {
  Reset()
  _ = NewCluster(1)
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
