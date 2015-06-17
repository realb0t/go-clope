package clope

import (
  "sync"
  "math"
  "github.com/realb0t/go-clope/io"
  clu "github.com/realb0t/go-clope/cluster"
  "github.com/realb0t/go-clope/cluster/store"
  tsn "github.com/realb0t/go-clope/transaction"
)

// Структура процесса
type Process struct {
  input io.Input
  output io.Output
  store store.ClusterStore
  r float64
}

// Создание нового процесса
func NewProcess(input io.Input, output io.Output, store store.ClusterStore, r float64) *Process {
  return &Process{input, output, store, r}
}

type SyncMsg struct {
  Delta float64
  Cluster *clu.Cluster
}

// Выбирает Лучший кластер или Cоздает Новый кластер,
// добавляет в него транзакцию и возвращает этот кластер
func (p *Process) BestClusterFor(t *tsn.Transaction) *clu.Cluster {
  var bestCluster *clu.Cluster

  if p.store.Len() > 0 {
    var wg sync.WaitGroup
    tempW := float64(len(t.Atoms))
    tempS := tempW
    deltaMax := tempS / math.Pow(tempW, p.r)
    syncDelta := make(chan *SyncMsg)

    wg.Add(p.store.Len())

    p.store.Iterate(func(c *clu.Cluster) {
      go func(cluster *clu.Cluster) {
        defer wg.Done()
        curDelta := cluster.DeltaAdd(t, p.r)
        syncDelta <- &SyncMsg{Delta: curDelta, Cluster: cluster}
      }(c)
    })

    go func() {
      wg.Wait() ; close(syncDelta)
    }()

    for msg := range syncDelta {
      if (msg.Delta > deltaMax) {
        deltaMax = msg.Delta
        bestCluster = msg.Cluster
      }
    }
  }

  if bestCluster == nil {
    bestCluster, _ = p.store.CreateCluster()
  }
  return bestCluster
}

// Инициализация первоначального размещения
func (p *Process) Initialization() {
  for trans := p.input.Pop(); trans != nil; trans = p.input.Pop() {
    bestCluster := p.BestClusterFor(trans)
    p.store.MoveTransaction(bestCluster.Id, trans)
    p.output.Push(trans)
  }
}

// Итерация по размещению с целью наилучшего
// расположения транзакций по кластерам
// За одну итерацию перемещается одна транзакция
func (p *Process) Iteration() {
  
  for {
    moved := false
    for trans := p.output.Pop(); trans != nil; trans = p.output.Pop() {
      lastClusterId := trans.ClusterId
      bestCluster := p.BestClusterFor(trans)
      if bestCluster.Id != lastClusterId {
        p.store.MoveTransaction(bestCluster.Id, trans)
        p.output.Push(trans)
        moved = true
      }
    }

    if !moved { 
      break
    }
  }
  _ = p.store.RemoveEmpty()
}

// Построение размещения с одной итерацией
func (p *Process) Build() {
  p.Initialization()
  p.Iteration()
}