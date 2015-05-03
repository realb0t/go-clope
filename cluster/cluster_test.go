package cluster

import "testing"

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