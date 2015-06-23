package clope

import (
  "log"
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
func (p *Process) BestClusterFor(t *tsn.Transaction) (*clu.Cluster, error) {
  var (
    bestCluster *clu.Cluster
    addError error
  )

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
    bestCluster, addError = p.store.CreateCluster()
  }
  return bestCluster, addError
}

// Инициализация первоначального размещения
func (p *Process) Initialization() error {
  var err error

  for trans := p.input.Pop(); trans != nil; trans = p.input.Pop() {
    bestCluster, err := p.BestClusterFor(trans)
    if err == nil {
      p.store.MoveTransaction(bestCluster.Id, trans)
    }
    p.output.Push(trans)

    if err != nil {
      break
    }
  }

  return err;
}

// Итерация по размещению с целью наилучшего
// расположения транзакций по кластерам
// За одну итерацию перемещается одна транзакция
func (p *Process) Iteration() {

  for {
    moved := false
    for trans := p.output.Pop(); trans != nil; trans = p.output.Pop() {
      lastClusterId := trans.ClusterId
      bestCluster, err := p.BestClusterFor(trans)
      
      if err != nil {
        panic(err)
      }

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

  if x := recover(); x != nil {
    log.Panicln(x)
  }

  _ = p.store.RemoveEmpty()
}

// Построение размещения с одной итерацией
func (p *Process) Build() {
  _ = p.Initialization()
  p.Iteration()
}