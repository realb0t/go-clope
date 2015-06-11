package cluster

import (
  "fmt"
  "math"
  "github.com/realb0t/go-clope/atom"
  trn "github.com/realb0t/go-clope/transaction"
)

// Структура кластера
type Cluster struct {
  Id int // ID
  // Кластерные характеристики
  N int // кол-во транзакций
  W int // кол-во уникальных объектов/атомов (или ширины кластера) 
  S int // площади кластера
  // Хеш атомов (элементов) кластера 
  // Ключ - это Атом
  // Значение - кол-во Атомов
  atoms map[*atom.Atom]int
}

// Создать новый кластер
func NewCluster(id int) *Cluster {
  return &Cluster{id, 0, 0, 0, make(map[*atom.Atom]int, 0)}
}

// Преобразование к строке для подробного вывода
func (c *Cluster) String() string {
  s := fmt.Sprintf("ID:%d;N:%d;W:%d;S:%d;%v", c.Id, c.N, c.W, c.S, c.atoms)
  return s
}

// Количество (частота) атомов в кластере
func (c *Cluster) Occ(atom *atom.Atom) int {
  return c.atoms[atom]
}

// Обновление хеша атомов кластера при добавлении транзакции
func (c *Cluster) RefreshAtomsAfterAdd(t *trn.Transaction) {
  for _, atom := range(t.Atoms) {
    if count, ok := c.atoms[atom] ; ok {
      c.atoms[atom] = count + 1
    } else {
      c.atoms[atom] = 1
    }
  }
  // c.refresh(c.N + 1)
}

// Обновление хеша атомов кластера при удалении транзакции
func (c *Cluster) RefreshAtomsAfterRemove(t *trn.Transaction) {
  for _, atom := range(t.Atoms) {
    if count, ok := c.atoms[atom] ; ok {
      c.atoms[atom] = count - 1
      if c.atoms[atom] == 0 {
        delete(c.atoms, atom)
      }
    }
  }
  // c.refresh(c.N - 1)
}

// Обновление кластерных характеристик
func (c *Cluster) Refresh(transCount int) {
  c.N = transCount
  c.W = len(c.atoms)
  c.S = 0

  for _, count := range(c.atoms) {
    c.S = c.S + count
  }
}

// Подсчет дельта-веса при добавлении транзакции в кластер
func (c *Cluster) DeltaAdd(t *trn.Transaction, r float64) float64 {
  transAtomsCount := len(t.Atoms)
  S_new := float64(c.S + transAtomsCount)
  W_new := float64(c.W)
  toCounter := transAtomsCount - 1
  for i := 0; i < toCounter; i++ {
    if _, ok := c.atoms[t.Atoms[i]]; !ok { W_new++; }
  }

  if c.N == 0 {
    return S_new / math.Pow(W_new, r)
  }

  profitCur := float64(c.S) * float64(c.N) / math.Pow(float64(c.W), r)
  profitNew := S_new * float64(c.N + 1) / math.Pow(W_new, r)
  return profitNew - profitCur
}
