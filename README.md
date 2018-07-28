# Chained

A Golang package for the quick creation of a simple blockchain.

## Basic Usage

```go
package main

import (
    "fmt"

    "github/rburmorrison/go-chained/pkg/chained"
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

## Installation

Run `go get github.com/rburmorrison/go-chained/...`

## Documentation

To view the Chained documentation, install the package and run `godoc -http=":6060"`. The documentation will be available at [http://localhost:6060/](http://localhost:6060/) within the packages page.

**Note:** $GOHOME must be set for this to work properly.