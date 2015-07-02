package atom

// Atoms it is map of strings

// Single atom
type Atom struct {
  Label string
}

// All atoms
var Atoms map[string]*Atom

// Reset all exist atoms
func Reset() {
  Atoms = make(map[string]*Atom, 0)
}

// Create and return new atom
// if this atom has exist then
// new atom don't create because return old
func NewAtom(label string) *Atom {
  if Atoms == nil { Reset() }
  if _, ok := Atoms[label]; !ok {
    Atoms[label] = &Atom{label}
  }
  return Atoms[label]
}

// Return atoms array by strings array
func NewAtoms(labels []string) []*Atom {
  atoms := make([]*Atom, 0)
  for _, label := range(labels) {
    atoms = append(atoms, NewAtom(label))
  }
  return atoms
}

// Atoms inspect
func (a *Atom) String() string {
  return a.Label
}