package clope

import (
  "sync"
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

type SyncMsg struct {
  Delta float64
  Cluster *clu.Cluster
}

// Выбирает Лучший кластер или Cоздает Новый кластер,
// добавляет в него транзакцию и возвращает этот кластер
func (p *Process) BestClusterFor(t *tsn.Transaction) *clu.Cluster {
  var bestCluster *clu.Cluster

  if len(clu.Clusters) > 0 {
    var wg sync.WaitGroup
    tempW := float64(len(t.Atoms))
    tempS := tempW
    deltaMax := tempS / math.Pow(tempW, p.r)
    syncDelta := make(chan *SyncMsg)

    wg.Add(len(clu.Clusters))

    for _, cluster := range(clu.Clusters) {
      go func(cluster *clu.Cluster) {
        defer wg.Done()
        curDelta := cluster.DeltaAdd(t, p.r)
        syncDelta <- &SyncMsg{Delta: curDelta, Cluster: cluster}
      }(cluster)
    }

    go func() {
      wg.Wait()
      close(syncDelta)
    }()

    for msg := range syncDelta {
      if (msg.Delta > deltaMax) {
        deltaMax = msg.Delta
        bestCluster = msg.Cluster
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
  
  for {
    moved := false

    for trans := p.output.Next(); trans != nil; trans = p.output.Next() {
      lastClusterId := trans.ClusterId
      bestCluster := p.BestClusterFor(trans)
      if bestCluster.Id != lastClusterId {
        bestCluster.MoveTransaction(trans)
        p.output.Write(trans)
        moved = true
      }
    }

    if !moved { break }
  }
  clu.RemoveEmpty()
}

// Построение размещения с одной итерацией
func (p *Process) Build() {
  clu.Reset()
  p.Initialization()
  p.Iteration()
}