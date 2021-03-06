package main

import "fmt"
import "os"

/*
	P(i) = \sum_{i%j==0}^i 10j

*/

func main() {
	for N := 100; ; N *= 10 {
		fmt.Printf("N = %v\n", N)
		ps := make([]int, N+1)
		for j := 1; j <= N; j++ {
			for i := j; i <= N; i += j {
				ps[i] += 10 * j
			}
		}
		for i, p := range ps {
			if p >= 34000000 {
				fmt.Printf("House %v is the first.\n", i)
				os.Exit(0)
			}
		}
	}
}
