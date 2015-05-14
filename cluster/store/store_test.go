package store_test

import (
  "testing"
  "github.com/realb0t/go-clope/cluster/store"
)

func TestNewProcess(t *testing.T) {
  _ = store.NewMemoryStore()
}

