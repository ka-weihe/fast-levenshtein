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

## Benchmarks
```bash
Benchmark/4/kaweihe-12         	   88651	     12858 ns/op	       0 B/op	       0 allocs/op
Benchmark/4/agniva-12          	   34550	     32099 ns/op	    7984 B/op	     499 allocs/op
Benchmark/4/arbovm-12          	   31778	     49311 ns/op	   23952 B/op	     499 allocs/op
Benchmark/4/dgryski-12         	   34489	     37951 ns/op	   23952 B/op	     499 allocs/op
Benchmark/8/kaweihe-12         	   51102	     22806 ns/op	       0 B/op	       0 allocs/op
Benchmark/8/agniva-12          	   17876	     66867 ns/op	   15968 B/op	     499 allocs/op
Benchmark/8/arbovm-12          	   12945	    104283 ns/op	   39920 B/op	     499 allocs/op
Benchmark/8/dgryski-12         	   11898	     92959 ns/op	   39920 B/op	     499 allocs/op
Benchmark/16/kaweihe-12        	   26749	     43723 ns/op	       0 B/op	       0 allocs/op
Benchmark/16/agniva-12         	    6129	    195921 ns/op	   23952 B/op	     499 allocs/op
Benchmark/16/arbovm-12         	    3370	    356006 ns/op	   71856 B/op	     499 allocs/op
Benchmark/16/dgryski-12        	    3242	    371193 ns/op	   71856 B/op	     499 allocs/op
Benchmark/32/kaweihe-12        	   12588	     93955 ns/op	       0 B/op	       0 allocs/op
Benchmark/32/agniva-12         	    1604	    767277 ns/op	   39920 B/op	     499 allocs/op
Benchmark/32/arbovm-12         	     838	   1434373 ns/op	  143712 B/op	     499 allocs/op
Benchmark/32/dgryski-12        	     837	   1427871 ns/op	  143712 B/op	     499 allocs/op
Benchmark/64/kaweihe-12        	    6538	    180181 ns/op	       0 B/op	       0 allocs/op
Benchmark/64/agniva-12         	     422	   2827182 ns/op	  327344 B/op	    1497 allocs/op
Benchmark/64/arbovm-12         	     219	   5449945 ns/op	  542912 B/op	    1497 allocs/op
Benchmark/64/dgryski-12        	     219	   5457779 ns/op	  542912 B/op	    1497 allocs/op
Benchmark/128/kaweihe-12       	     856	   1447393 ns/op	  526944 B/op	    1996 allocs/op
Benchmark/128/agniva-12        	     100	  10688841 ns/op	  654688 B/op	    1497 allocs/op
Benchmark/128/arbovm-12        	      55	  21351738 ns/op	 1085824 B/op	    1497 allocs/op
Benchmark/128/dgryski-12       	      56	  21179082 ns/op	 1085824 B/op	    1497 allocs/op
Benchmark/256/kaweihe-12       	     242	   4947683 ns/op	 1053888 B/op	    1996 allocs/op
Benchmark/256/agniva-12        	      28	  41670764 ns/op	 1309376 B/op	    1497 allocs/op
Benchmark/256/arbovm-12        	      13	  85361092 ns/op	 2171648 B/op	    1497 allocs/op
Benchmark/256/dgryski-12       	      13	  84091060 ns/op	 2171648 B/op	    1497 allocs/op
Benchmark/512/kaweihe-12       	      69	  17196304 ns/op	 2107776 B/op	    1996 allocs/op
Benchmark/512/agniva-12        	       7	 164796786 ns/op	 2618752 B/op	    1497 allocs/op
Benchmark/512/arbovm-12        	       3	 346575811 ns/op	 4471040 B/op	    1497 allocs/op
Benchmark/512/dgryski-12       	       3	 338171785 ns/op	 4471040 B/op	    1497 allocs/op
Benchmark/1024/kaweihe-12      	      18	  65227427 ns/op	 4215552 B/op	    1996 allocs/op
Benchmark/1024/agniva-12       	       2	 650856037 ns/op	 5237504 B/op	    1497 allocs/op
Benchmark/1024/arbovm-12       	       1	1388587828 ns/op	 8814336 B/op	    1497 allocs/op
Benchmark/1024/dgryski-12      	       1	1351116629 ns/op	 8814336 B/op	    1497 allocs/op
```

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
