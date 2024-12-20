# go-singleton

A simple Go library that implement singleton

## Usage

```go
package main

import (
	"fmt"
	"github.com/Wing924/go-singleton"
)


var instance singleton.Singleton[string]

func GetInstance() string {
    return instance.GetOrInit(func() string {
        return "Hello, World!"
    })
}

func main() {
    fmt.Println(GetInstance())
    // Output: Hello, World!
}

```
