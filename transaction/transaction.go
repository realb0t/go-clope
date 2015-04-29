package transaction

import (
  "github.com/realb0t/go-clope/atom"
  cl "github.com/realb0t/go-clope/cluster"
)

type Transaction struct {
  Items []*atom.Atom // Атомы транзакции
  Cluster *cl.Cluster // Текущий кластер транзакции
}

// Создание новой транзакции из массива атомов
func NewTransaction(atoms []*atom.Atom) *Transaction {
  return &Transaction{atoms}
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
