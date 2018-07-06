# Chained

A Golang package for the quick creation of a simple blockchain.

## Basic Usage

```go
package main

import (
    "fmt"

    "ryanburmeister/chained"
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

To install chained, clone the repository into your `$GOHOME/src` folder. On Unix machines, $GOHOME tends to resolve to `$HOME/go`. Once cloned, ensure the folder name is "chained", navigate to the directory in a terminal and run `go install`. Chained is built entirely with the standard packages, so no third-party packages are required.

## Documentation

To view the Chained documentation, install the package and run `godoc -http=":6060"`. The documentation will be available at [http://localhost:6060/](http://localhost:6060/) within the packages page.

**Note:** $GOHOME must be set for this to work properly.