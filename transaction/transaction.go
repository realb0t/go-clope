package transaction

import (
  "github.com/realb0t/go-clope/atom"
)

// Структура транзакции
type Transaction struct {
  Atoms []*atom.Atom // Атомы транзакции
  ClusterId int
}

// Создание новой транзакции из массива атомов
func NewTransaction(atoms []*atom.Atom) *Transaction {
  return &Transaction{atoms, -1}
}

// Создание новой транзации из массива строк или атомов
func BuildTransaction(atoms interface{}) *Transaction {
  var trans *Transaction
  switch a := atoms.(type) {
    case []string:
      trans = NewTransaction(atom.NewAtoms(a))
    case []*atom.Atom:
      trans = NewTransaction(a)
    default:
      panic("Not allowed transaction atoms type")
  }
  return trans
}

func (t *Transaction) String() []*atom.Atom {
  return t.Atoms
}