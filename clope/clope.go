package clope

import (
  "math"
  "github.com/realb0t/go-clope/io"
  clu "github.com/realb0t/go-clope/cluster"
  tsn "github.com/realb0t/go-clope/transaction"
)

type Process struct {
  reader *io.Reader
  writer *io.Writer
  r float64
  clusters []*clu.Cluster
}

func NewProcess(reader *io.Reader, writer *io.Writer, r float64) {
  return &Process{reader, writer, r, make([]*clu.Cluster, 0)}
}

func (p *Process) Clusters() []*clu.Cluster {
  return clusters;
}

// Создает новый кластер и добавляет в него транзакцию
func (p *Process) addToNewCluster(t *tsn.Transaction) *clu.Cluster {
    newCluster = cluster.NewCluster()
    newCluster.AddTransaction(t)
    p.clusters = append(p.clusters, newCluster)
    bestCluster = newCluster
}

// Выбирает или создает лучший кластер, добавляет
// в него транзакцию и возвращает этот кластер
func (p *Process) bestClusterFor(t *tsn.Transaction) *clu.Cluster {
  var bestCluster *clu.Cluster

  if len(p.clusters) > 0 {
    tempW := float64(count(t.Items))
    tempS := tempW
    deltaMax = tempS / Math.pow(tempW, p.r)

    for _, cluster := range(p.clusters) {
      curDelta = cluster.DeltaAdd(t)
      if (curDelta > deltaMax) {
        deltaMax = curDelta
        bestCluster = cluster
      }
    }
  }

  if bestCluster == nil {
    bestCluster = p.addToNewCluster(t)
  }

  return bestCluster
}

func (p *Process) Initialization() {
  for trans := p.reader.Next() {
    cluster := p.bestClusterFor(trans)
    p.writer.Write(trans, cluster)
  }
}

func (p *Process) Iteration() {
  moved := false
  for moved == false {
    for trans := p.reader.Next() {
      cluster := p.bestClusterFor(trans)
      if cluster.Id != trans.Cluster.Id {
        p.writer.Write(trans, cluster)
        moved = true
      }
    }
  }
}

func (p *Process) Build() {
  p.Initialization()
  p.Iteration()
}