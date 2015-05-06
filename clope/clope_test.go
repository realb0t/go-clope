package clope

import (
  "testing"
  "github.com/realb0t/go-clope/io"
  a "github.com/realb0t/go-clope/atom"
  tr "github.com/realb0t/go-clope/transaction"
  _ "github.com/realb0t/go-clope/cluster"
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
    tr.NewTransaction(a.NewAtoms([]string{ "a", "b" })),
    tr.NewTransaction(a.NewAtoms([]string{ "a", "b", "c" })),
    tr.NewTransaction(a.NewAtoms([]string{ "a", "c", "d" })),
    tr.NewTransaction(a.NewAtoms([]string{ "d", "e" })),
    tr.NewTransaction(a.NewAtoms([]string{ "d", "e", "f" })),
    tr.NewTransaction(a.NewAtoms([]string{ "h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d" })),
    tr.NewTransaction(a.NewAtoms([]string{ "h", "e", "l", "l", "o" })),
    tr.NewTransaction(a.NewAtoms([]string{ "w", "o", "r", "l", "d" })),
  }

  input   := io.NewMemoryInput(&trans)
  output  := io.NewMemoryOutput()
  process := NewProcess(input, output, 2.6)
  process.Build()
  //cluster.Print()
}
