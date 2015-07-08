package simple

import (
  "testing"
  "github.com/realb0t/go-clope/atom"
)

func TestNewSimpleTransaction(t *testing.T) {
  _ = NewSimple(make([]*atom.Atom, 0))
}

func TestUniqTransactionAtoms(t *testing.T) {
  trans := MakeUniq("a", "b", "b", "c", "c", "c")
  if len(trans.Atoms()) != 3 {
    t.Error("Uncorrect transaction atoms count")
  }
}

func TestBuildTransaction(t *testing.T) {
  trans1 := BuildTransaction([]string{ "a", "b", "c" })
  if len(trans1.Atoms()) != 3 {
    t.Error("Not correct atoms for first transaction")
  }

  trans2 := BuildTransaction([]string{ "a", "b", "c", "d" })
  if len(trans2.Atoms()) != 4 {
    t.Error("Not correct atoms for second transaction")
  }
}