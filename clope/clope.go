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
  clusters map[*clu.Cluster]bool
}

// Создание нового процесса
func NewProcess(reader *io.Reader, writer *io.Writer, r float64) {
  return &Process{reader, writer, r, make([]*clu.Cluster, 0)}
}

// Возвращает массив кластеров
func (p *Process) Clusters() []*clu.Cluster {
  clusters := make([]*clu.Cluster, 0)
  for cluster, _ := range(p.clusters) {
    clusters = append(clusters, cluster)
  }
  return clusters;
}

// Создает новый кластер и добавляет в него транзакцию
func (p *Process) CreateCluster(t *tsn.Transaction) *clu.Cluster {
    newClasterId := len(p.clusters) + 1
    newCluster := cluster.NewCluster()
    return newCluster
}

// Выбирает Лучший кластер или Cоздает Новый кластер,
// добавляет в него транзакцию и возвращает этот кластер
func (p *Process) BestClusterFor(t *tsn.Transaction) *clu.Cluster {
  var bestCluster *clu.Cluster

  if len(p.clusters) > 0 {
    tempW := float64(count(t.Items))
    tempS := tempW
    deltaMax := tempS / Math.pow(tempW, p.r)

    // Эту часть алгоритма возможно распараллелить
    // если потребуется работа с большим количеством
    // кластеров
    for cluster, _ := range(p.clusters) {
      curDelta = cluster.DeltaAdd(t, p.r)
      if (curDelta > deltaMax) {
        deltaMax = curDelta
        bestCluster = cluster
      }
    }
  }

  if bestCluster == nil {
    bestCluster = p.CreateCluster(t)
  }
  return bestCluster
}

// Инициализация первоначального размещения
func (p *Process) Initialization() {
  for trans := p.reader.Next() {
    bestCluster := p.BestClusterFor(trans)
    bestCluster.MoveTransaction(trans)
    p.clusters[bestCluster] = true
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
      lastCluster := trans.Cluster
      bestCluster := p.BestClusterFor(trans)
      // Исли "лучший" кластер не текущий кластер
      if bestCluster.Id != lastCluster.Id {
        bestCluster.MoveTransaction(trans)
        p.clusters[bestCluster] = true
        if lastCluster.N == 0 {
          delete(p.clusters, lastCluster)
        }
        p.writer.Write(trans)
        moved = true
      }
    }
  }
}

// Построение размещения с одной итерацией
func (p *Process) Build() {
  p.Initialization()
  p.Iteration()
}