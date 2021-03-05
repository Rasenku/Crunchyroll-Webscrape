package main

import (
	"testing"
)


func Benchmarkmain(b *testing.B) {
    for i := 0; i < b.N; i++ {
      main()
    }
}
