package memory

import (
  "testing"
  a "github.com/realb0t/go-clope/atom"
  tr "github.com/realb0t/go-clope/transaction"
)

func TestNewMemoryInput(t *testing.T) {
  _ = NewMemoryInput(&[]*tr.Transaction{})
}

func TestNewMemoryOutput(t *testing.T) {
  _ = NewMemoryOutput()
}

func TestMemoryInputPop(t *testing.T) {
  trans := []*tr.Transaction{ 
    tr.NewTransaction(a.NewAtoms([]string{ "a", "b" })),
    tr.NewTransaction(a.NewAtoms([]string{ "c" })),
  }
  input := NewMemoryInput(&trans)
  firstTrans, _ := input.Pop()
  secondTrans, _ := input.Pop()
  thirdTrans, _ := input.Pop()

  if firstTrans != trans[1] {
    t.Error("Pop transaction", firstTrans, " != ", trans[1] )
  }

  if secondTrans != trans[0] {
    t.Error("Pop transaction", secondTrans, " != ", trans[0] )
  }

  if thirdTrans != nil {
    t.Error("Pop transaction", secondTrans, " != nil")
  }
}

func TestMemoryOutputPush(t *testing.T) {
  trans := []*tr.Transaction{ 
    tr.NewTransaction(a.NewAtoms([]string{ "a", "b" })),
    tr.NewTransaction(a.NewAtoms([]string{ "c" })),
  }
  output := NewMemoryOutput()
  res, _ := output.Pop()
  if res != nil {
    t.Error("Not empty first Pop value")
  }

  _ = output.Push(trans[0])
  _ = output.Push(trans[1])

  if len(output.Data) != 2 {
    t.Error("Not correct data size")
  }

  fPopTran, _  := output.Pop()
  if fPopTran != trans[1] {
    t.Error("Pop transaction", fPopTran, " != ", trans[1] )
  }

  sPopTran, _ := output.Pop()
  if sPopTran != trans[0] {
    t.Error("Pop transaction", sPopTran, " != ", trans[0] )
  }

  _ = output.Push(fPopTran)
  _ = output.Push(sPopTran)

  if len(output.Data) != 2 {
    t.Error("Not correct data size")
  }
}