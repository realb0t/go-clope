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
// то не создает новый атом а возвращает старый
func NewAtom(label string) *Atom {
  if Atoms == nil { Reset() }
  if _, ok := Atoms[label]; !ok {
    Atoms[label] = &Atom{label}
  }
  return Atoms[label]
}

// Возвращает массив атомов в
// сооответствии с массивом строк
func NewAtoms(labels []string) []*Atom {
  atoms := make([]*Atom, 0)
  for _, label := range(labels) {
    atoms = append(atoms, NewAtom(label))
  }
  return atoms
}

func (a *Atom) String() string {
  return a.Label
}