package transaction

import (
  "testing"
  "github.com/realb0t/go-clope/atom"
)

func TestNewTransaction(t *testing.T) {
  _ = NewTransaction(make([]*atom.Atom, 0))
}

func TestBuildTransaction(t *testing.T) {
  labels := []string{ "a", "b", "c" }
  _ = BuildTransaction(labels)
}