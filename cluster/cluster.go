package cluster

import (
  "math"
  "github.com/realb0t/go-clope/atom"
  trans "github.com/realb0t/go-clope/transaction"
)

// Структура кластера
type Cluster struct {
  Id int // ID
  transactions []*trans.Transaction // массив транзакций
  // Кластерные характеристики
  N int // количества транзакций
  W int // числа уникальных объектов (или ширины кластера) 
  S int // площади кластера
  atoms map[*atom.Atom]int // хеш атомов (элементов) кластера 
}

// Создать новый кластер
func NewCluster(id int) {
  return &Cluster{id, [], 0, 0, 0, []}
}

// Количество (частота) атомов в кластере
func (c *Cluster) Occ(atom *atom.Atom) int {
  return с.atoms[atom]
}

// Обновление хеша атомов кластера при добавлении транзакции
func (c *Cluster) RefreshAtomsAfterAdd(t *trans.Transaction) {
  for _, atom := range(t.Items) {
    if c.atoms[atom] == nil {
      c.atoms[atom] := 1
    } else {
      c.atoms[atom] = c.atoms[atom] + 1
    }
  }
}

// Обновление хеша атомов кластера при удалении транзакции
func (c *Cluster) RefreshAtomsAfterRemove(t *trans.Transaction) {
  for _, atom := range(t.Items) {
    if c.atoms[atom] != nil {
      c.atoms[atom] = c.atoms[atom] - 1
      if c.atoms[atom] == 0 {
        delete(c.atoms, atom)
      }
    }
  }
}

// Обновление кластерных характеристик
func (c *Cluster) Refresh() {
  c.N = len(c.transactions)
  c.W = len(c.atoms)
  c.S = 0

  for _, a := range(c.atoms) {
    c.S = c.S + c.Occ(a)
  }
}

// Добавление транзакции в кластер
func (c *Cluster) MoveTransaction(t *trans.Transaction) {
  // Добавляем транзакцию к текущему кластеру
  c.transactions = append(c.transactions, t)
  // Обновляем кластерные характеристики
  c.RefreshAtomsAfterAdd(t)
  c.Refresh()
  // Удаляем транзакцию из старого кластера (если есть)
  if t.Cluster != nil {
    t.Cluster.RemoveTransaction(t)
  }
  // и переключаем указатель кластера в транзакции
  // на новый.
  t.Cluster = c
}

// Удаление транзакции из кластера
func (c *Cluster) RemoveTransaction(t *trans.Transaction) {
  ei := -1

  for i, trans := range(c.transactions) {
    if (t == trans) { ei = i }
  }

  copy(c.transactions[ei:], c.transactions[ei+1:])
  c.transactions[len(c.transactions)-1] = nil
  c.transactions = c.transactions[:len(c.transactions)-1]

  c.RefreshAtomsAfterRemove()
  c.Refresh()
}

// Подсчет дельта-веса при добавлении транзакции в кластер
func (c *Cluster) DeltaAdd(t *Transaction, r float64) float64 {
  tItemsCount := len(t.Items)
  S_new := float64(c.S + tItemsCount)
  W_new := c.W
  toCounter := tItemsCount - 1
  for i := 0; i < toCounter; i++ {
    if c.Occ(t.Items[i]) == nil { W_new++; }
  }

  if c.N == 0 {
    return S_new / Math.pow(W_new, r)
  }

  profitCur := c.S * c.N / Math.pow(c.W, r)
  profitNew := S_new * (c.N + 1) / Math.pow(W_new, r)
  return profitNew - profitCur
}
