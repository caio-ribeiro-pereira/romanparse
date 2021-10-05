[![Build Status](https://app.travis-ci.com/caio-ribeiro-pereira/romanparse.svg?branch=main)](https://app.travis-ci.com/caio-ribeiro-pereira/romanparse)
[![Coverage](https://gocover.io/_badge/github.com/caio-ribeiro-pereira/romanparse)](https://gocover.io/github.com/caio-ribeiro-pereira/romanparse)
[![Go Report Card](https://goreportcard.com/badge/github.com/caio-ribeiro-pereira/romanparse)](https://goreportcard.com/report/github.com/caio-ribeiro-pereira/romanparse)
[![Godoc Reference](https://godoc.org/github.com/caio-ribeiro-pereira/romanparse?status.svg)](http://godoc.org/github.com/caio-ribeiro-pereira/romanparse)

# romanparse

Simple go package which converts roman strings to integer

## Example

```go
import (
  "fmt"
  "github.com/caio-ribeiro-pereira/romanparse"
)

func main() {
  number, err := romanparse.ToInteger("XVI")
  if err == nil {
    fmt.Print(number) // 16
  }
}
```

## Install

```bash
go get -u github.com/caio-ribeiro-pereira/romanparse
```