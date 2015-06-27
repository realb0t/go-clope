package clope

import (
  "errors"
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
  store *store.Store
  r float64
}

// Создание нового процесса
func NewProcess(input io.Input, output io.Output, str *store.Store, r float64) *Process {
  return &Process{input, output, str, r}
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
  var (
    err error
    trans *tsn.Transaction
    bestCluster *clu.Cluster
  )

  for {
    trans, err = p.input.Pop()
    if err != nil { break }
    if trans == nil { break }

    bestCluster, err = p.BestClusterFor(trans)
    if err != nil { break }

    p.store.MoveTransaction(bestCluster.Id, trans)
    err = p.output.Push(trans)
    if err != nil { break }
  }

  return err
}

// Итерация по размещению с целью наилучшего
// расположения транзакций по кластерам
// За одну итерацию перемещается одна транзакция
func (p *Process) Iteration() (returnError error) {
  defer func() {
    if err := recover(); err != nil {
      log.Panicf("Iteration Error: %v\n", err)
      returnError = errors.New( "Iteration error" )
    }
  }()

  for {
    moved := false

    var (
      trans *tsn.Transaction
      err error
      bestCluster *clu.Cluster
    )

    for {
      trans, err = p.output.Pop()
      if err != nil { panic(err) }
      if trans == nil { break }

      lastClusterId := trans.ClusterId
      bestCluster, err = p.BestClusterFor(trans)
      if err != nil { panic(err) }

      if bestCluster.Id != lastClusterId {
        p.store.MoveTransaction(bestCluster.Id, trans)
        err = p.output.Push(trans)
        if err != nil { panic(err) }
        moved = true
      }
    }

    if !moved { break }
  }

  err := p.store.RemoveEmpty()
  return err
}

// Построение размещения с одной итерацией
func (p *Process) Build() error {
  err := p.Initialization()
  if err != nil {
    log.Panicf("Initialization Error: %v\n", err)
    return err 
  }
  return p.Iteration()
}