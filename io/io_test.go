package io

import (
  "testing"
  a "github.com/realb0t/go-clope/atom"
  tr "github.com/realb0t/go-clope/transaction"
)

func TestNewMemoryInput(t *testing.T) {
  _ = NewMemoryInput(&[]*tr.Transaction{})
}

func TestNewMemoryOutput(t *testing.T) {
  _ = NewMemoryOutput()
}

func TestMemoryInputNext(t *testing.T) {
  trans := []*tr.Transaction{ 
    tr.NewTransaction(a.NewAtoms([]string{ "a", "b" })),
    tr.NewTransaction(a.NewAtoms([]string{ "c" })),
  }
  input := NewMemoryInput(&trans)
  nextTran := input.Next()
  if nextTran != trans[1] {
    t.Error("Next transaction", nextTran, " != ", trans[1] )
  }
}

func TestMemoryOutputWrite(t *testing.T) {
  trans := []*tr.Transaction{ 
    tr.NewTransaction(a.NewAtoms([]string{ "a", "b" })),
    tr.NewTransaction(a.NewAtoms([]string{ "c" })),
  }
  output := NewMemoryOutput()

  if output.Next() != nil {
    t.Error("Not empty first Next value")
  }

  output.Write(trans[0])
  output.Write(trans[1])

  if len(output.Data) != 2 {
    t.Error("Not correct data size")
  }

  fNextTran := output.Next()
  if fNextTran != trans[1] {
    t.Error("Next transaction", fNextTran, " != ", trans[1] )
  }

  sNextTran := output.Next()
  if sNextTran != trans[0] {
    t.Error("Next transaction", sNextTran, " != ", trans[0] )
  }

  output.Write(fNextTran)
  output.Write(sNextTran)

  if len(output.Data) != 2 {
    t.Error("Not correct data size")
  }
}