package atom

type Atom struct {
  Label string
}

var Atoms map[string]*Atom

// Сбросить все существующие атомы
func Reset() {
  Atoms = make(map[string]*Atom, 0)
}

// Создание нового атома
// если такой атом уже существует
// то операция возвращает nil
func NewAtom(label string) *Atom {
  if Atoms == nil {
    Reset()
  }

  if Atoms[label] != nil {
    return nil
  } else {
    Atoms[label] = &Atom{label}
    return Atoms[label]
  }
}

func NewAtoms(labels []string) []*Atom {
  atoms := make([]*Atom, 0)
  for _, label := range(labels) {
    newAtom := NewAtom(label)
    if newAtom != nil {
      atoms = append(atoms, newAtom)
    }
  }
  return atoms
}