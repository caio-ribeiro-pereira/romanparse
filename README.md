# romanparse

Simple go package which converts roman strings to integer

## Example

```go
import (
  "fmt"
  "github.com/caio-ribeiro-pereira/romanparse"
)

func main() {
  roman := "XVI"
  fmt.Printf("%s = %i\n", roman, romanparse.ToInteger(roman))
  // XVI = 16
}
```

## Install

```bash
go get -u github.com/caio-ribeiro-pereira/romanparse
```