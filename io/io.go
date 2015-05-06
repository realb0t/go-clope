package io

import (
  "github.com/realb0t/go-clope/transaction"
)

// Интерфейс Ввода
type Input interface {
  Next() *transaction.Transaction
}

// Интерфейс Вывода
type Output interface {
  Input
  Write(*transaction.Transaction)
}

// Тестовый Ввода (test input)
type MemoryInput struct {
  data []*transaction.Transaction
}

// Создать новый тестовый Ввод
func NewMemoryInput(data []*transaction.Transaction) *MemoryInput {
  return &MemoryInput{data}
}

// Тестовый Вывод (test output)
type MemoryOutput struct {
  Data []*transaction.Transaction
}

// Создать новый тестовый Вывод
func NewMemoryOutput() *MemoryOutput {
  return &MemoryOutput{[]*transaction.Transaction{}}
}

// Извлечь следующее значение из Ввода
func (r *MemoryInput) Next() *transaction.Transaction {
  var trans *transaction.Transaction
  trans, r.data = r.data[len(r.data)-1], r.data[:len(r.data)-1]
  return trans
}

// Извлечь следующее значение из Вывода
func (r *MemoryOutput) Next() *transaction.Transaction {
  if len(r.Data) > 0 {
    var trans *transaction.Transaction
    trans, r.Data = r.Data[len(r.Data)-1], r.Data[:len(r.Data)-1]
    return trans
  } else {
    return nil
  }
}

// Записать значение в ввод и сбросить счетчик извлечения значений
func (r *MemoryOutput) Write(trans *transaction.Transaction) {
  r.Data = append(r.Data, trans)
}