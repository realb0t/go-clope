package clope

import (
  "math"
  "github.com/realb0t/go-clope/io"
  clu "github.com/realb0t/go-clope/cluster"
  tsn "github.com/realb0t/go-clope/transaction"
)

// Структура процесса
type Process struct {
  reader io.Reader
  writer io.Writer
  r float64
}

// Создание нового процесса
func NewProcess(reader io.Reader, writer io.Writer, r float64) *Process {
  return &Process{reader, writer, r}
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
  for trans := p.reader.Next(); trans != nil; {
    bestCluster := p.BestClusterFor(trans)
    bestCluster.MoveTransaction(trans)
    p.writer.Write(trans)
    trans = p.reader.Next()
  }
}

// Итерация по размещению с целью наилучшего
// расположения транзакций по кластерам
// За одну итерацию перемещается одна транзакция
func (p *Process) Iteration() {
  moved := false
  for moved == false {
    for trans := p.reader.Next(); trans != nil; {
      lastClusterId := trans.ClusterId
      // Ищем наилучший клстер для данной транзакции
      bestCluster := p.BestClusterFor(trans)
      // Eсли "лучший" кластер не текущий кластер
      if bestCluster.Id != lastClusterId {
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