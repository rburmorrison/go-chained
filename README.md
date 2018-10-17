# Chained

A Golang package for the quick creation of a simple blockchain.

**Documentation:** [https://godoc.org/github.com/rburmorrison/go-chained](https://godoc.org/github.com/rburmorrison/go-chained)

**Please note:** This package is meant to be used for prototyping, experimentation, or education only. Some key cryptography is missing from this package to make it fully secure. However, it is great for getting your ideas made quickly.

**Install:** `go get github.com/rburmorrison/go-chained`

## Basic Usage

```go
package main

import (
    "fmt"

    "github/rburmorrison/go-chained"
)

func main() {
    // Set difficulty (Must be done before creation)
    chained.Target = 2

    // Create a new blockchain
    b := chained.NewBlockchain()

    // Mine new block (Will take time)
    b.MineNewBlockAndApply()

    // View blockchain JSON
    fmt.Println(b.JSONString())
}
```
