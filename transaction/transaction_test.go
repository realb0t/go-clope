package transaction

import (
  "testing"
  "github.com/realb0t/go-clope/atom"
)

func TestNewTransaction(t *testing.T) {
  _ = NewTransaction(make([]*atom.Atom, 0))
}

func TestBuildTransaction(t *testing.T) {
  trans1 := BuildTransaction([]string{ "a", "b", "c" })
  if len(trans1.Atoms) != 3 {
    t.Error("Not correct atoms for first transaction")
  }

  trans2 := BuildTransaction([]string{ "a", "b", "c", "d" })
  if len(trans2.Atoms) != 4 {
    t.Error("Not correct atoms for second transaction")
  }
}