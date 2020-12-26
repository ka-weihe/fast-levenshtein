# fast-levenshtein :rocket: 

> Fastest levenshtein implementation in Go.
> Measure the difference between two strings.

```bash
$ go get github.com/ka-weihe/fast-levenshtein
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/ka-weihe/fast-levenshtein"
)

func main() {
	s1 := "fast"
	s2 := "fastest"
	distance := levenshtein.Distance(s1, s2)
	fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
	// => The distance between fast and fastest is 3.
}
```

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
