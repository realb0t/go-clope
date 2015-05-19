package io

import (
  "github.com/realb0t/go-clope/transaction"
)

// Интерфейс Ввода
type Input interface {
  // Pop next transaction in output/output
  Pop() *transaction.Transaction
}

// Интерфейс Вывода
type Output interface {
  Input
  // Add transaction in output
  Push(*transaction.Transaction)
}

// Тестовый Ввода (test input)
type MemoryInput struct {
  data []*transaction.Transaction
}

// Создать новый тестовый Ввод
func NewMemoryInput(data *[]*transaction.Transaction) *MemoryInput {
  return &MemoryInput{*data}
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
func (r *MemoryInput) Pop() *transaction.Transaction {
  lenData := len(r.data)
  if lenData == 0 { return nil }
  var trans *transaction.Transaction
  trans, r.data = r.data[len(r.data)-1], r.data[:len(r.data)-1]
  return trans
}

// Извлечь следующее значение из Вывода
func (r *MemoryOutput) Pop() *transaction.Transaction {
  lenData := len(r.Data)
  if lenData == 0 { return nil }
  var trans *transaction.Transaction
  trans, r.Data = r.Data[lenData-1], r.Data[:lenData-1]
  return trans
}

// Записать значение в ввод и сбросить счетчик извлечения значений
func (r *MemoryOutput) Push(trans *transaction.Transaction) {
  r.Data = append(r.Data, trans)
}

