package clope

import (
  "testing"
  "github.com/realb0t/go-clope/io"
  a "github.com/realb0t/go-clope/atom"
  tr "github.com/realb0t/go-clope/transaction"
  cl "github.com/realb0t/go-clope/cluster"
)

func TestNewProcess(t *testing.T) {
  trans := []*tr.Transaction{
    tr.NewTransaction(a.NewAtoms([]string{ "a" })),
  }
  _ = NewProcess(io.NewMemoryInput(&trans),
    io.NewMemoryOutput(), 1.0)
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
  process := NewProcess(input, output, 1.8)
  process.Build()
  
  clusterCheck := (
    cl.Clusters[1].Tran(0) == trans[7] &&
    cl.Clusters[1].Tran(1) == trans[6] &&
    cl.Clusters[1].Tran(2) == trans[5] &&
    cl.Clusters[2].Tran(0) == trans[2] &&
    cl.Clusters[2].Tran(1) == trans[1] &&
    cl.Clusters[2].Tran(2) == trans[0] &&
    cl.Clusters[2].Tran(3) == trans[3] &&
    cl.Clusters[2].Tran(4) == trans[4])

  if !clusterCheck {
    cl.Print()
    t.Error("Not valid clusters")
  }
}
