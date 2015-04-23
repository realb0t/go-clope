package clope

import (
  "github.com/realb0t/go-clope/io"
  "github.com/realb0t/go-clope/cluster"
)

type Process struct {
  reader *io.Reader
  writer *io.Writer
  clusters []*cluster.Cluster
}

func NewProcess(reader *io.Reader, writer *io.Writer) {
  return &Process{reader, writer, make([]*cluster.Cluster, 0)}
}

func (p *Process) Clusters() []*cluster.Cluster {
  return clusters;
}

func (p *Process) bestClusterFor(t *transaction.Transaction) *cluster.Cluster {
}

func (p *Process) Initialization() {
  for trans := p.reader.Next() {
    cluster := p.bestClusterFor(trans)
    p.writer.Write(trans, cluster)
  }
}

func (p *Process) Iteration() {

}

func (p *Process) Build() {

}