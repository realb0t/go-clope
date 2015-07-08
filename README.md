# go-clope

Golang implementation of [CLOPE](https://www.google.ru/search?q=clope) clusterization algorythm with parallel calculation.

## Install and Run example

Install go-clope
```
$ go get github.com/realb0t/go-clope
```

Example programm:
```go
package clope_example

import (
  "github.com/realb0t/go-clope/clope"
  "github.com/realb0t/go-clope/cluster/store"
  "github.com/realb0t/go-clope/transaction"
  "github.com/realb0t/go-clope/transaction/simple"
  io "github.com/realb0t/go-clope/io/memory"
  driver "github.com/realb0t/go-clope/cluster/store/driver/memory"
)

func main() {
  trans := []*transaction.Transaction{ 
    simple.Make( "a", "b" ),
    simple.Make( "a", "b", "c" ),
    simple.Make( "a", "c", "d" ),
    simple.Make( "d", "e" ),
    simple.Make( "d", "e", "f" ),
    simple.Make( "h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d" ),
    simple.Make( "h", "e", "l", "l", "o" ),
    simple.Make( "w", "o", "r", "l", "d" ),
  }

  repulsion := 1.8
  input     := io.NewMemoryInput(&trans)
  output    := io.NewMemoryOutput()
  driver    := driver.NewMemory()
  storage   := store.NewStore(driver)
  process   := clope.NewProcess(input, output, store, repulsion)
  err       := process.Build()

  if (err != nil) {
    panic(err)
  }

  // Print created clusters
  storage.Print()
}
```

Output:
```
[1] - [[w o r l d] [h e l l o] [h e l l o   w o r l d]]
[2] - [[a c d] [a b c] [a b] [d e] [d e f]]
```

Output Pop/Push can include into DB-transaction.

## Data Structures

**Transaction** - is data-transaction structure have atoms collection.

**Atom** - is data structure for build transaction entity (as string). 
Example: `gender:female`, `income:site-from.com`, etc.

**Clusters** - map of all linked transaction.

**Input** - is data-store source for transactions

**Output** - is data-store for work algorythm. Before session is empty and after session is empty.

**Store** - is data-store for storage results as array of clusters with linked transactions. 

## Usage

Package `go-clope/clope` have struct `clope.Process` with three methods:
`Initialization`, `Iteration` and `Build`.

Method `Initialization` build first clusters structure by transactions from **Input**. And Push linked transactions into **Output**.

Method `Iteration` transactions distributes between clusters. And Push transactions into **Output**.

Method `Build` call clusters reset and call `Initialization` and `Iteration`.

Base usage:
```go
  process := clope.NewProcess(input, output, repulsion)
  err := process.Build()
```

After `Build` you can call again `Iteration`.

For read results you can make use of your IO Output or `cluster.Clusters` map or `cluster.Print()` function for print.

## IO interface

Package `go-clope/io` have test IO-structures `MemoryInput` and `MemoryOutput` for memory data storing . 

You can create other IO structures for other data stores (PostgreSQL, MongoDB, Redis, etc.).

Your IO structures must implement for interface:
```go
type Input interface {
  // Pop Pop transaction unlinked from data store.
  Pop() (*transaction.Transaction, error)
}

type Output interface {
  // Pop Pop linked (to cluster) transaction. After initialization process.
  Input
  // Push linked (to cluster) transaction into data store.
  Push(*transaction.Transaction) error
}
```

## Get Results

Result storage is Cluster store and then driver.

For get Clusters as `map[int]*cluster.Cluster`:
```go
  clusters, errors := storage.Driver().Clusters()
```

For iterate Clusters:
```go
  storage.Driver().Iterate(func(clu *cluster.Cluster) {
    // ...
  })
```

For acces for Cluster Transaction use Driver.ClusterTransactions function:
```go
  storage.Driver().Iterate(func(clu *cluster.Cluster) {
    transactions := storage.Driver().ClusterTransactions(clu)
  })
```

Or get all transaction as `map[int][]*transaction.Transaction`:
```go
  transactions, error := storage.Driver().Transactions()
```

## TODO

- Abstract Transaction (with custom fields)
- Apply decimal calculations
- Create clustarization testing tools