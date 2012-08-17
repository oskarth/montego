package main

import (
	"math/rand"
  "os"
  "runtime/pprof"
  "time"
  "fmt"
)

// BENCHMARKS
//
// n = 1 000 000
// GOMAXPROCS 1: 1.15s no, 600ms
// GOMAXPROCS 4: 4+s, no, 240ms!

const NCPU = 4

func withinCircle(x float64, y float64) bool {
	return ((x*x + y*y) <= 1)
}

func worker(n int, ch chan float64) {
	var num int
  source := rand.NewSource(time.Now().UnixNano())
  random := rand.New(source)
  for i := 0; i < n; i++ {
    if withinCircle(random.Float64(), random.Float64()) {
			num++
		}
	}

  ch <- float64(num) / float64(n)
}

func main() {
  f, err := os.Create("montego.pprof")
  if err != nil {
    panic(err)
  }
  if err := pprof.StartCPUProfile(f); err != nil {
    panic(err)
  }
  defer pprof.StopCPUProfile()

  ch := make(chan float64, NCPU)
	n := 10000000 // number of points total
  nNCPU := n / NCPU
	var pi float64

  t0 := time.Now()
	for i := 0; i < NCPU; i++ {
		go worker(nNCPU, ch)
	}

	for i := 0; i < NCPU; i++ {
		pi += <-ch
	}

  pi = (pi / float64(NCPU)) * 4.0

  t1 := time.Now()
	fmt.Println(pi, "took ", t1.Sub(t0), "to calculate.")
}
