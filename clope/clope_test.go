package clope

import (
  "testing"
  "github.com/realb0t/go-clope/atom"
  "github.com/realb0t/go-clope/cluster/store"
  "github.com/realb0t/go-clope/transaction"
  "github.com/realb0t/go-clope/transaction/simple"
  io "github.com/realb0t/go-clope/io/memory"
  driver "github.com/realb0t/go-clope/cluster/store/driver/memory"
)

func TestNewProcess(t *testing.T) {
  trans := []transaction.Transaction{
    simple.NewSimple(atom.NewAtoms([]string{ "a" })),
  }

  driver := driver.NewMemory()
  store  := store.NewStore(driver)
  input  := io.NewMemoryInput(&trans)
  output := io.NewMemoryOutput()

  _ = NewProcess(input, output, store, 1.0)
}

func TestBuildIntegration(t *testing.T) {
  trans := []transaction.Transaction{ 
    simple.Make( "a", "b" ),
    simple.Make( "a", "b", "c" ),
    simple.Make( "a", "c", "d" ),
    simple.Make( "d", "e" ),
    simple.Make( "d", "e", "f" ),
    simple.Make( "h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d" ),
    simple.Make( "h", "e", "l", "l", "o" ),
    simple.Make( "w", "o", "r", "l", "d" ),
  }

  input   := io.NewMemoryInput(&trans)
  output  := io.NewMemoryOutput()
  driver  := driver.NewMemory()
  storage := store.NewStore(driver)
  process := NewProcess(input, output, storage, 1.8)
  _        = process.Build()

  clusterTransactions, _ := storage.Driver().Transactions()
  clusterCheck := (
    clusterTransactions[1][0] == trans[7] &&
    clusterTransactions[1][1] == trans[6] &&
    clusterTransactions[1][2] == trans[5] &&
    clusterTransactions[2][0] == trans[2] &&
    clusterTransactions[2][1] == trans[1] &&
    clusterTransactions[2][2] == trans[0] &&
    clusterTransactions[2][3] == trans[3] &&
    clusterTransactions[2][4] == trans[4])

  if !clusterCheck {
    storage.Print()
    t.Error("Not valid clusters")
  }
}

func TestWithOtherOrders(t *testing.T) {
  trans := []transaction.Transaction{ 
    simple.Make( "a", "b" ),
    simple.Make( "b", "a" ),
    simple.Make( "c", "d" ),
    simple.Make( "d", "c", "b" ),
  }

  input   := io.NewMemoryInput(&trans)
  output  := io.NewMemoryOutput()
  driver  := driver.NewMemory()
  storage := store.NewStore(driver)
  process := NewProcess(input, output, storage, 3.0)
  process.Build()

  clusterTransactions, _ := storage.Driver().Transactions()
  clusterCheck := (
    clusterTransactions[1][0] == trans[3] &&
    clusterTransactions[1][1] == trans[2] &&
    clusterTransactions[2][0] == trans[0] &&
    clusterTransactions[2][1] == trans[1])

  if !clusterCheck {
    storage.Print()
    t.Error("Not valid clusters")
  }
}

func TestWithUniqTransactions(t *testing.T) {
  t.SkipNow()
  // @todo Unskip after create testing tools

  trans := []transaction.Transaction{ 
    simple.MakeUniq( "a", "b", "a", "b", "c", "c" ),
    simple.MakeUniq( "c", "b", "b", "c", "a" ),
    simple.MakeUniq( "c", "d", "d", "c", "c", "d" ),
    simple.MakeUniq( "d", "c", "b", "d", "b" ),
  }

  input   := io.NewMemoryInput(&trans)
  output  := io.NewMemoryOutput()
  driver  := driver.NewMemory()
  storage := store.NewStore(driver)
  process := NewProcess(input, output, storage, 3.325)
  process.Build()

  storage.Print()
}
