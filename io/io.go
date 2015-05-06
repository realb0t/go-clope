package io

import (
  "github.com/realb0t/go-clope/transaction"
)

type Reader interface {
  Next() *transaction.Transaction
}

type Writer interface {
  Write(*transaction.Transaction)
}
