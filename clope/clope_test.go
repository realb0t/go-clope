package clope

import (
  "testing"
  "github.com/realb0t/go-clope/io"
  a "github.com/realb0t/go-clope/atom"
  tr "github.com/realb0t/go-clope/transaction"
)

func TestNewProcess(t *testing.T) {
  trans := []*tr.Transaction{
    tr.NewTransaction(a.NewAtoms([]string{ "a" })),
  }
  _ = NewProcess(io.NewMemoryInput(&trans),
    io.NewMemoryOutput(), 1.0)
}