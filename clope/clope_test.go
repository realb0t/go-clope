package clope

import (
  "testing"
  "github.com/realb0t/go-clope/io"
  a "github.com/realb0t/go-clope/atom"
  tr "github.com/realb0t/go-clope/transaction"
  "github.com/realb0t/go-clope/cluster/store"
)

func TestNewProcess(t *testing.T) {
  trans := []*tr.Transaction{
    tr.NewTransaction(a.NewAtoms([]string{ "a" })),
  }
  _ = NewProcess(io.NewMemoryInput(&trans),
    io.NewMemoryOutput(), store.NewMemoryStore(), 1.0)
}

func TestBuildIntegration(t *testing.T) {
  trans := []*tr.Transaction{ 
    tr.Make( "a", "b" ),
    tr.Make( "a", "b", "c" ),
    tr.Make( "a", "c", "d" ),
    tr.Make( "d", "e" ),
    tr.Make( "d", "e", "f" ),
    tr.Make( "h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d" ),
    tr.Make( "h", "e", "l", "l", "o" ),
    tr.Make( "w", "o", "r", "l", "d" ),
  }

  input   := io.NewMemoryInput(&trans)
  output  := io.NewMemoryOutput()
  storage := store.NewMemoryStore()
  process := NewProcess(input, output, storage, 1.8)
  process.Build()

  clusterTransactions := make(map[int][]*tr.Transaction, 0)
  clusters, _ := storage.Clusters()
  for cId, cluster := range(clusters) {
    clusterTransactions[cId], _ = storage.ClusterTransactions(cluster)
  }

  clusterCheck := (
    clusterTransactions[1][0] == trans[7] &&
    clusterTransactions[1][1] == trans[6] &&
    clusterTransactions[1][2] == trans[5] &&
    clusterTransactions[2][0] == trans[2] &&
    clusterTransactions[2][1] == trans[1] &&
    clusterTransactions[2][2] == trans[0] &&
    clusterTransactions[2][3] == trans[3] &&
    clusterTransactions[2][4] == trans[4])

  if !clusterCheck {
    storage.Print()
    t.Error("Not valid clusters")
  }
}

func TestWithOtherOrders(t *testing.T) {
  trans := []*tr.Transaction{ 
    tr.Make( "a", "b" ),
    tr.Make( "b", "a" ),
    tr.Make( "c", "d" ),
    tr.Make( "d", "c", "b" ),
  }

  input   := io.NewMemoryInput(&trans)
  output  := io.NewMemoryOutput()
  storage := store.NewMemoryStore()
  process := NewProcess(input, output, storage, 3.0)
  process.Build()

  clusterTransactions := make(map[int][]*tr.Transaction, 0)
  clusters, _ := storage.Clusters()
  for cId, cluster := range(clusters) {
    clusterTransactions[cId], _ = storage.ClusterTransactions(cluster)
  }

  clusterCheck := (
    clusterTransactions[1][0] == trans[3] &&
    clusterTransactions[1][1] == trans[2] &&
    clusterTransactions[2][0] == trans[0] &&
    clusterTransactions[2][1] == trans[1])

  if !clusterCheck {
    storage.Print()
    t.Error("Not valid clusters")
  }
}

func TestWithUniqTransactions(t *testing.T) {
  t.SkipNow()
  // @todo Unskip after create testing tools

  trans := []*tr.Transaction{ 
    tr.MakeUniq( "a", "b", "a", "b", "c", "c" ),
    tr.MakeUniq( "c", "b", "b", "c", "a" ),
    tr.MakeUniq( "c", "d", "d", "c", "c", "d" ),
    tr.MakeUniq( "d", "c", "b", "d", "b" ),
  }

  input   := io.NewMemoryInput(&trans)
  output  := io.NewMemoryOutput()
  storage := store.NewMemoryStore()
  process := NewProcess(input, output, storage, 3.325)
  process.Build()

  storage.Print()
}
