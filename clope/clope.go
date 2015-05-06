package clope

import (
  "math"
  "github.com/realb0t/go-clope/io"
  clu "github.com/realb0t/go-clope/cluster"
  tsn "github.com/realb0t/go-clope/transaction"
)

// Структура процесса
type Process struct {
  input io.Input
  output io.Output
  r float64
}

// Создание нового процесса
func NewProcess(input io.Input, output io.Output, r float64) *Process {
  return &Process{input, output, r}
}

// Выбирает Лучший кластер или Cоздает Новый кластер,
// добавляет в него транзакцию и возвращает этот кластер
func (p *Process) BestClusterFor(t *tsn.Transaction) *clu.Cluster {
  var bestCluster *clu.Cluster

  if len(clu.Clusters) > 0 {
    tempW := float64(len(t.Atoms))
    tempS := tempW
    deltaMax := tempS / math.Pow(tempW, p.r)

    for _, cluster := range(clu.Clusters) {
      // Эту часть программы возможно распараллелить
      // если потребуется работа с большим количеством
      // кластеров
      curDelta := cluster.DeltaAdd(t, p.r)
      if (curDelta > deltaMax) {
        deltaMax = curDelta
        bestCluster = cluster
      }
    }
  }

  if bestCluster == nil {
    bestCluster = clu.AddCluster()
  }
  return bestCluster
}

// Инициализация первоначального размещения
func (p *Process) Initialization() {
  for trans := p.input.Next(); trans != nil; trans = p.input.Next() {
    bestCluster := p.BestClusterFor(trans)
    bestCluster.MoveTransaction(trans)
    p.output.Write(trans)
  }
}

// Итерация по размещению с целью наилучшего
// расположения транзакций по кластерам
// За одну итерацию перемещается одна транзакция
func (p *Process) Iteration() {
  moved := false
  for moved == false {
    for trans := p.output.Next(); trans != nil; trans = p.output.Next() {
      lastClusterId := trans.ClusterId
      bestCluster := p.BestClusterFor(trans)
      if bestCluster.Id != lastClusterId {
        bestCluster.MoveTransaction(trans)
        p.output.Write(trans)
        moved = true
      }
    }
  }
  clu.RemoveEmpty()
}

// Построение размещения с одной итерацией
func (p *Process) Build() {
  clu.Reset()
  p.Initialization()
  p.Iteration()
}