package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func fibonacci(iterations uint64) []uint64 {
	var results []uint64
	for n := uint64(0); n < iterations; n++ {
		if n == 0 {
			results = append(results, 0)
		}
		if n == 1 {
			results = append(results, 1)
		}
		if n > 1 {
			f := (results[len(results)-1]) + (results[len(results)-2])
			results = append(results, f)
		}
	}
	return results
}

func compression(size int) {
	var b bytes.Buffer
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = byte(rand.Intn(1337))
	}
	w := gzip.NewWriter(&b)
	w.Write(data)
	w.Close()
	r, err := gzip.NewReader(&b)
	if err != nil {
		panic(err)
	}
	r.Close()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Number of logical cores:", runtime.NumCPU())
	t0 := time.Now()
	fibonacci(100000000)
	t1 := time.Now()
	fmt.Printf("Fibonacci took %v to run.\n", t1.Sub(t0))
	t0 = time.Now()
	compression(10000000)
	t1 = time.Now()
	fmt.Printf("gzip took %v to run.\n", t1.Sub(t0))
}
