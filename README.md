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
  "github.com/realb0t/go-clope/io"
  "github.com/realb0t/go-clope/transaction"
  "github.com/realb0t/go-clope/cluster"
  "github.com/realb0t/go-clope/cluster/store"
)

func main() {
  trans := []*transaction.Transaction{ 
    tr.Make( "a", "b" ),
    tr.Make( "a", "b", "c" ),
    tr.Make( "a", "c", "d" ),
    tr.Make( "d", "e" ),
    tr.Make( "d", "e", "f" ),
    tr.Make( "h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d" ),
    tr.Make( "h", "e", "l", "l", "o" ),
    tr.Make( "w", "o", "r", "l", "d" ),
  }

  repulsion := 1.8
  input     := io.NewMemoryInput(&trans)
  output    := io.NewMemoryOutput()
  storage   := store.MemoryStore()
  process   := clope.NewProcess(input, output, store, repulsion)
  process.Build()

  // All created clusters in map[*atom.Atom]int 
  // cluster.Clusters

  // All transaction with clusters put to IO.Output.Push

  // Print created clusters
  store.Print()
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
  process.Build()
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
  Pop() *transaction.Transaction
}

type Output interface {
  // Pop Pop linked (to cluster) transaction. After initialization process.
  Input
  // Push linked (to cluster) transaction into data store.
  Push(*transaction.Transaction)
}
```

## Result store interface



## TODO

- Storage interface methods with errors
- Add Store/Drivers for Cluster/Store for abstract mechanics of storage
- Abstract Transaction (with custom fields)
- Apply decimal calculations
- Create clustarization testing tools