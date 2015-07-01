package transaction

import (
  "fmt"
  "github.com/realb0t/go-clope/atom"
)

// Структура транзакции
type Transaction interface {
  Atoms() []*atom.Atom
  String() string
  GetClusterId() int
  SetClusterId(int) Transaction
}

// Структура транзакции
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

// Создает новую непривязанную транзакцию
func NewSimpleTransaction(atoms []*atom.Atom) Transaction {
  return &SimpleTransaction{atoms, -1}
}

// Создает транзакцию по набору String парамеров
func Make(stringAtoms ...string) Transaction {
  return NewSimpleTransaction(atom.NewAtoms(stringAtoms))
}

// Создает транзакцию с уникальным набором атомов
// по набору String парамеров
func MakeUniq(stringAtoms ...string) Transaction {
  return NewUniqTransaction(atom.NewAtoms(stringAtoms))
}

// Создание новой транзакции из массива атомов
// новая транзакция не имеет состоит из уникальных атомов
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

