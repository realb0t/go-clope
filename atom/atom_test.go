package atom

import "testing"

func TestNewAtom(t *testing.T) {
  Reset()
  _ = NewAtom("a")
  if len(Atoms) != 1 {
    t.Error("Not add atoms")
  }
}

func TestNewAtoms(t *testing.T) {
  Reset()
  labels := []string{ "a", "b", "b" }
  atoms := NewAtoms(labels)
  if len(atoms) != 2 {
    t.Error("Atoms dublicated", Atoms, len(atoms))
  }
}