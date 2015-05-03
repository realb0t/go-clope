package cluster

import (
  "math"
  "github.com/realb0t/go-clope/atom"
  trn "github.com/realb0t/go-clope/transaction"
)

// Структура кластера
type Cluster struct {
  Id int // ID
  // @todo Возможно тут должна быть Map
  // чтобы нельзя было поместить в один кластер несколько
  // одинаковых транзакций
  transactions []*trn.Transaction // массив транзакций
  // Кластерные характеристики
  N int // количества транзакций
  W int // числа уникальных объектов (или ширины кластера) 
  S int // площади кластера
  // Хеш атомов (элементов) кластера 
  // Ключ - это Атом
  // Значение - кол-во Атомов
  atoms map[*atom.Atom]int 
}

// Созданные кластеры
var Clusters map[int]*Cluster

// Сбросить набор кластеров
func Reset() {
  Clusters = make(map[int]*Cluster, 0)
}

// Создать новый кластер
func NewCluster(id int) *Cluster {
  return &Cluster{id, make([]*trn.Transaction, 0), 0, 0, 0, make(map[*atom.Atom]int, 0)}
}

// Создать и Добавить новый кластер
func AddCluster() *Cluster {
  if Clusters == nil { Reset() }
  nextId := len(Clusters)
  Clusters[nextId] = NewCluster(nextId)
  return Clusters[nextId]
}

// Количество (частота) атомов в кластере
func (c *Cluster) Occ(atom *atom.Atom) int {
  return c.atoms[atom]
}

// Обновление хеша атомов кластера при добавлении транзакции
func (c *Cluster) refreshAtomsAfterAdd(t *trn.Transaction) {
  for _, atom := range(t.Atoms) {
    if count, ok := c.atoms[atom] ; ok {
      c.atoms[atom] = count + 1
    } else {
      c.atoms[atom] = 1
    }
  }
}

// Обновление хеша атомов кластера при удалении транзакции
func (c *Cluster) refreshAtomsAfterRemove(t *trn.Transaction) {
  for _, atom := range(t.Atoms) {
    if count, ok := c.atoms[atom] ; ok {
      c.atoms[atom] = count - 1
      if c.atoms[atom] == 0 {
        delete(c.atoms, atom)
      }
    }
  }
}

// Обновление кластерных характеристик
func (c *Cluster) refresh() {
  c.N = len(c.transactions)
  c.W = len(c.atoms)
  c.S = 0

  for _, count := range(c.atoms) {
    c.S = c.S + count
  }
}

// Добавление/Перемещение транзакции в кластер
func (c *Cluster) MoveTransaction(t *trn.Transaction) {
  // Добавляем транзакцию к текущему кластеру
  c.transactions = append(c.transactions, t)
  // Обновляем кластерные характеристики
  c.refreshAtomsAfterAdd(t)
  c.refresh()
  // Если для транзакции был определен кластер
  if t.ClusterId != -1 {
    // Удаляем транзакцию из старого кластера
    Clusters[t.ClusterId].removeTransaction(t)
  }
  // и переключаем указатель кластера в транзакции
  // на текущий кластер
  t.ClusterId = c.Id
}

// Удаление транзакции из кластера
func (c *Cluster) removeTransaction(t *trn.Transaction) {
  ei := -1

  // Опеределяем индекс данной транзакции
  for i, trans := range(c.transactions) {
    if (t == trans) { ei = i }
  }

  // Если данная транзакция определенна в массиве
  if ei != -1 {
    // Извлекаем ее из массива транзакций
    copy(c.transactions[ei:], c.transactions[ei+1:])
    c.transactions[len(c.transactions)-1] = nil
    c.transactions = c.transactions[:len(c.transactions)-1]

    c.refreshAtomsAfterRemove(t)
    c.refresh()
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
