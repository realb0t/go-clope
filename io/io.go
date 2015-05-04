package io

import (
  "github.com/realb0t/go-clope/transaction"
)

type Input interface {
  Next() *transaction.Transaction
}

type Output interface {
  Input
  Write(*transaction.Transaction)
}

type MemoryInput struct {
  data []*transaction.Transaction
}

type MemoryOutput struct {
  Data []*transaction.Transaction
  index int
}

func (r *MemoryInput) Next() *transaction.Transaction {
  var trans *transaction.Transaction
  trans, r.data = r.data[len(r.data)-1], r.data[:len(r.data)-1]
  return trans
}

func (r *MemoryOutput) Next() *transaction.Transaction {
  if r.index > 0 {
    r.index--
    var trans *transaction.Transaction
    trans, r.Data = r.Data[len(r.Data)-1], r.Data[:len(r.Data)-1]
    return trans
  } else {
    return nil
  }
}

func (r *MemoryOutput) Write(trans *transaction.Transaction) {
  r.Data = append(r.Data, trans)
  r.index = len(r.Data)
}