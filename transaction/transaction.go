package transaction

import (
  "fmt"
  "github.com/realb0t/go-clope/atom"
)

// Transaction interface
type Transaction interface {
  // Return atoms array
  Atoms() []*atom.Atom
  // Inspect
  String() string
  // Get current cluster id
  GetClusterId() int
  // Set transaction cluster id
  SetClusterId(int) Transaction
}

// Transaction without addition attributes
// for tests
type SimpleTransaction struct {
  atoms []*atom.Atom // Атомы транзакции
  clusterId int
}

func (t *SimpleTransaction) Atoms() []*atom.Atom {
  return t.atoms
}

func (t *SimpleTransaction) String() string {
  return fmt.Sprintf("%v", t.atoms)
}

func (t *SimpleTransaction) GetClusterId() int {
  return t.clusterId
}

func (t *SimpleTransaction) SetClusterId(clusterId int) Transaction {
  t.clusterId = clusterId
  return t
}

// Create new simple transaction
func NewSimpleTransaction(atoms []*atom.Atom) Transaction {
  return &SimpleTransaction{atoms, -1}
}

// Construct transactions by strings parameters
func Make(stringAtoms ...string) Transaction {
  return NewSimpleTransaction(atom.NewAtoms(stringAtoms))
}

// Construct transactions with uniq atoms by strings parameters
func MakeUniq(stringAtoms ...string) Transaction {
  return NewUniqTransaction(atom.NewAtoms(stringAtoms))
}

// Create new simple transaction with uniq atoms
func NewUniqTransaction(atoms []*atom.Atom) Transaction {
  uniqAtoms := make([]*atom.Atom, 0)
  atomsMap := make(map[*atom.Atom]bool, 0)
  for _, atom := range(atoms) { atomsMap[atom] = true }
  for atom, _ := range(atomsMap) { uniqAtoms = append(uniqAtoms, atom) }
  return NewSimpleTransaction(uniqAtoms)
}

// Создание новой транзации из массива строк или атомов
func BuildTransaction(atoms interface{}) Transaction {
  var trans Transaction
  switch a := atoms.(type) {
    case []string:
      trans = NewSimpleTransaction(atom.NewAtoms(a))
    case []*atom.Atom:
      trans = NewSimpleTransaction(a)
    default:
      panic("Not allowed transaction atoms type")
  }
  return trans
}

