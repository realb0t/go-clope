package transaction

import (
  "github.com/realb0t/go-clope/atom"
)

// Transaction interface
type Transaction interface {
  // Return atoms array
  Atoms() []*atom.Atom
  // Inspect
  String() string
  // Get current cluster id
  GetClusterId() int
  // Set transaction cluster id
  SetClusterId(int) Transaction
}



