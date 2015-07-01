package io

import (
  "github.com/realb0t/go-clope/transaction"
)

// Интерфейс Ввода
// Очереди не обработанных объектов
type Input interface {
  // Pop next transaction in output/output
  Pop() (transaction.Transaction, error)
}

// Интерфейс Вывода
// Очереди объектов на обработку
type Output interface {
  Input
  // Add transaction in output
  Push(transaction.Transaction) error
}
