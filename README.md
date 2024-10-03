# Go Snowflake ID Generator

This is a simple implementation of the Snowflake algorithm in Go, which generates unique, 64-bit IDs.

## Features
- Time-based unique ID generation
- Supports up to 1024 nodes
- Up to 4096 unique IDs per millisecond per node

## Installation

To use this package, you can run:

```bash
go get github.com/yourusername/go-snowflake
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/tony-zhuo/snowflake-go"
)

func main() {
	node, err := snowflake.NewSnowflake(1)
	if err != nil {
		fmt.Println("Error creating node:", err)
		return
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Generated ID:", node.NextID())
	}
}
```