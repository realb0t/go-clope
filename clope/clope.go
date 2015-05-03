package clope

import (
  "math"
  "github.com/realb0t/go-clope/io"
  clu "github.com/realb0t/go-clope/cluster"
  tsn "github.com/realb0t/go-clope/transaction"
)

// Структура процесса
type Process struct {
  reader *io.Reader
  writer *io.Writer
  r float64
}

// Создание нового процесса
func NewProcess(reader *io.Reader, writer *io.Writer, r float64) {
  return &Process{reader, writer, r, make([]*clu.Cluster, 0)}
}

// Выбирает Лучший кластер или Cоздает Новый кластер,
// добавляет в него транзакцию и возвращает этот кластер
func (p *Process) BestClusterFor(t *tsn.Transaction) *clu.Cluster {
  var bestCluster *clu.Cluster

  if len(clu.Clusters) > 0 {
    tempW := float64(count(t.Atoms))
    tempS := tempW
    deltaMax := tempS / Math.pow(tempW, p.r)

    // Эту часть алгоритма возможно распараллелить
    // если потребуется работа с большим количеством
    // кластеров
    for id, cluster := range(clu.Clusters) {
      curDelta = cluster.DeltaAdd(t, p.r)
      if (curDelta > deltaMax) {
        deltaMax = curDelta
        bestCluster = cluster
      }
    }
  }

  if bestCluster == nil {
    bestCluster = cluster.AddCluster()
  }
  return bestCluster
}

// Инициализация первоначального размещения
func (p *Process) Initialization() {
  for trans := p.reader.Next() {
    bestCluster := p.BestClusterFor(trans)
    bestCluster.MoveTransaction(trans)
    p.writer.Write(trans, bestCluster)
  }
}

// Итерация по размещению с целью наилучшего
// расположения транзакций по кластерам
// За одну итерацию перемещается одна транзакция
func (p *Process) Iteration() {
  moved := false
  for moved == false {
    for trans := p.reader.Next() {
      // @todo Избавиться от .Cluster
      lastCluster := trans.Cluster
      bestCluster := p.BestClusterFor(trans)
      // Исли "лучший" кластер не текущий кластер
      if bestCluster.Id != lastCluster.Id {
        bestCluster.MoveTransaction(trans)
        p.writer.Write(trans)
        moved = true
      }
    }
  }
}

// Построение размещения с одной итерацией
func (p *Process) Build() {
  clu.Reset()
  p.Initialization()
  p.Iteration()
}