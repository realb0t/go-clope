package cluster_test

import (
  "testing"
  "github.com/realb0t/go-clope/cluster"
  "github.com/realb0t/go-clope/cluster/store"
  trn "github.com/realb0t/go-clope/transaction"
  drv "github.com/realb0t/go-clope/cluster/store/driver/memory"
)

func TestNewCluster(t *testing.T) {
  _ = cluster.NewCluster(1)
}

func TestDeltaAddEvaluative(t *testing.T) {
  r := 2.6
  d := drv.NewMemory()
  s := store.NewStore(d)

  trans := []trn.Transaction{
    trn.Make( "a", "b" ),
    trn.Make( "a", "b", "c" ),
    trn.Make( "a", "c", "d" ),
    trn.Make( "d", "e" ),
    trn.Make( "d", "e", "f" ),
  }

  clusters := make([]*cluster.Cluster, 4)
  clusters[0], _ = s.CreateCluster()
  clusters[1], _ = s.CreateCluster()
  clusters[2], _ = s.CreateCluster()
  clusters[3], _ = s.CreateCluster()

  s.MoveTransaction(clusters[0].Id, trans[0])
  s.MoveTransaction(clusters[0].Id, trans[1])
  s.MoveTransaction(clusters[0].Id, trans[2])
  s.MoveTransaction(clusters[1].Id, trans[3])
  s.MoveTransaction(clusters[1].Id, trans[4])
  s.MoveTransaction(clusters[2].Id, trans[0])
  s.MoveTransaction(clusters[2].Id, trans[1])
  s.MoveTransaction(clusters[3].Id, trans[2])
  s.MoveTransaction(clusters[3].Id, trans[3])
  s.MoveTransaction(clusters[3].Id, trans[4])

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
