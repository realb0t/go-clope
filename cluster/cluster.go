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
  // @todo Возможно тут должна быть Map
  // чтобы нельзя было поместить в один кластер несколько
  // одинаковых транзакций
  // @todo По возможности убрать хранение транзакций
  // в кластере, т.к. нужно нормализовать хранение данных
  // и при работе программы и исключить их дублирование
  // (дублирование транзакций в кластере)
  transactions []*trn.Transaction // массив транзакций
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
  return &Cluster{id, make([]*trn.Transaction, 0), 0, 0, 0, make(map[*atom.Atom]int, 0)}
}

// Пустой ли кластер
func (c *Cluster) IsEmpty() bool {
  return len(c.transactions) == 0
}

// Get transaction by index
func (c *Cluster) GetTransaction(i int) *trn.Transaction {
  return c.transactions[i]
}

// Get all cluster transaction
func (c *Cluster) GetTransactions() []*trn.Transaction {
  return c.transactions
}

// Преобразование к строке
func (c *Cluster) String() string {
  s := fmt.Sprintf("[%d] - %v", c.Id, c.transactions)
  return s
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

// Удаление транзакции из кластера
func (c *Cluster) RemoveTransaction(t *trn.Transaction) {
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

func (c *Cluster) AddTransaction(t *trn.Transaction) {
  // Добавляем транзакцию к текущему кластеру
  c.transactions = append(c.transactions, t)
  // Обновляем кластерные характеристики
  c.refreshAtomsAfterAdd(t)
  c.refresh()
  // и переключаем указатель кластера в транзакции
  // на текущий кластер
  t.ClusterId = c.Id
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
