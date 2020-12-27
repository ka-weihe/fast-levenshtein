package levenshtein_test

import (
	"fmt"
	agnivade "github.com/agnivade/levenshtein"
	arbovm "github.com/arbovm/levenshtein"
	dgryski "github.com/dgryski/trifles/leven"
	kaweihe "github.com/ka-weihe/fast-levenshtein"
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func rnd_string(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func rnd_string_arr(arr_len int, str_len int) []string {
	b := make([]string, arr_len)
	for i := range b {
		b[i] = rnd_string(str_len)
	}
	return b
}

func TestFuzz(t *testing.T) {
	for i := 0; i < 100000; i++ {
		str1 := rnd_string(rand.Intn(500))
		str2 := rnd_string(rand.Intn(500))
		re1 := kaweihe.Distance(str1, str2)
		re2 := agnivade.ComputeDistance(str1, str2)
		if re1 != re2 {
			t.Errorf("Test[%d]: Distance(%q,%q)", i, re1, re2)
		}
	}
}

func Benchmark(b *testing.B) {
	data := [...][]string{
		rnd_string_arr(500, 4),
		rnd_string_arr(500, 8),
		rnd_string_arr(500, 16),
		rnd_string_arr(500, 32),
		rnd_string_arr(500, 64),
		rnd_string_arr(500, 128),
		rnd_string_arr(500, 256),
		rnd_string_arr(500, 512),
		rnd_string_arr(500, 1024),
	}
	tmp := 0
	for i, pick := range data {
		b.Run(fmt.Sprint(1<<(i+2)), func(b *testing.B) {
			b.Run("kaweihe", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					for n := 0; n < len(pick)-1; n++ {
						tmp += kaweihe.Distance(pick[n], pick[n+1])
					}
				}
			})
			b.Run("agniva", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					for n := 0; n < len(pick)-1; n++ {
						tmp += agnivade.ComputeDistance(pick[n], pick[n+1])
					}
				}
			})
			b.Run("arbovm", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					for n := 0; n < len(pick)-1; n++ {
						tmp += arbovm.Distance(pick[n], pick[n+1])
					}
				}
			})
			b.Run("dgryski", func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					for n := 0; n < len(pick)-1; n++ {
						tmp += dgryski.Levenshtein([]rune(pick[n]), []rune(pick[n+1]))
					}
				}
			})
		})
	}
	println(tmp)
}
