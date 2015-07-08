package simple

import (
  "fmt"
  "github.com/realb0t/go-clope/atom"
  "github.com/realb0t/go-clope/transaction"
)

// Transaction without addition attributes
// for tests
type Simple struct {
  atoms []*atom.Atom // Атомы транзакции
  clusterId int
}

func (t *Simple) Atoms() []*atom.Atom {
  return t.atoms
}

func (t *Simple) String() string {
  return fmt.Sprintf("%v", t.atoms)
}

func (t *Simple) GetClusterId() int {
  return t.clusterId
}

func (t *Simple) SetClusterId(clusterId int) transaction.Transaction {
  t.clusterId = clusterId
  return t
}

// Create new simple transaction
func NewSimple(atoms []*atom.Atom) transaction.Transaction {
  return &Simple{atoms, -1}
}

// Construct transactions by strings parameters
func Make(stringAtoms ...string) transaction.Transaction {
  return NewSimple(atom.NewAtoms(stringAtoms))
}

// Construct transactions with uniq atoms by strings parameters
func MakeUniq(stringAtoms ...string) transaction.Transaction {
  return NewUniqTransaction(atom.NewAtoms(stringAtoms))
}

// Create new simple transaction with uniq atoms
func NewUniqTransaction(atoms []*atom.Atom) transaction.Transaction {
  uniqAtoms   := make([]*atom.Atom, 0)
  atomsMap    := make(map[*atom.Atom]bool, 0)
  for _, atom := range(atoms) { atomsMap[atom] = true }
  for atom, _ := range(atomsMap) { uniqAtoms = append(uniqAtoms, atom) }
  return NewSimple(uniqAtoms)
}

// Создание новой транзации из массива строк или атомов
func BuildTransaction(atoms interface{}) transaction.Transaction {
  var trans transaction.Transaction
  switch a := atoms.(type) {
    case []string:
      trans = NewSimple(atom.NewAtoms(a))
    case []*atom.Atom:
      trans = NewSimple(a)
    default:
      panic("Not allowed transaction atoms type")
  }
  return trans
}