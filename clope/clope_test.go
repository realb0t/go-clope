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
  process := NewProcess(input, output, 2.6)
  process.Build()
  //cluster.Print()
}
