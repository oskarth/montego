package main

import (
  "math"
  "math/rand"
  "time"
  "fmt"
)

// TERMINOLOGY
// 
// UNIT OF WORK - a function which approximates PI by doing 1000 withinCircles
// BATCH - A collection of n UNITS. Each unit is spawned as a go routine.
// ERRTERM - max error between delta(max,min)
// BATCHES - list of three BATCHes.
// DELTA - batches min vs batches max

// FLOW
//
// in main function we have a inf for loop which breaks when delta < errterm
// in for loop we run three batches
// each batch keep spawning go routines to do units
// these go routines send their results to batch-number-CHAN
// batch-number-CHAN is a "list" of all units of work
// ... i am sure something is missing here-ish :D
// we read from batch-number-CHAN and add all those unit results to batch[n] we 
// we calc average of batch[n] for n=1,2,3 and put that into our BATCHES list
// we check DELTA, else repeat

func withinCircle(x float64, y float64) bool {
  return ((x * x + y * y) <= 1)
}

// Approximates Pi. Launched as a seperate process.
func calcPi (n int, c chan float64) { // wait how does this work?
  // I mean, how do we pass in the seed so rand.Float64() can use it
  circleArea := 0.0
  totalPoints := 0.0
  for i := 0; i < n; i++ {
    var x = rand.Float64()
    var y = rand.Float64()

    if (withinCircle(x,y)) {
      circleArea += 1.0
    }
    totalPoints += 1.0
  }
  c <- 4.0 * (circleArea / totalPoints) // send approximation to channel
}

func main() {
 rand.Seed(time.Now().Unix())

  epsilon := 1e-6
  iterations := 3
  values := make([]float64, iterations, iterations)
  piChannel := make(chan float64)

  for { // breaks when we reach the termination condition
    for i := 0; i < iterations; i++ {
      go calcPi(1000, piChannel)
    }
    for i := 0; i < iterations; i++ {
      values[i] = <-piChannel // reads from channel and puts in values list
    }

    // TODO: Un-lame
    mymax := math.Max(math.Max(values[0], values[1]), values[2])
    mymin := math.Min(math.Min(values[0], values[1]), values[2])

    delta := math.Abs(mymax - mymin)
    if delta < epsilon { break }
  }
  fmt.Println(values[0])
  fmt.Println("Delta real pi our pi ", math.Abs(math.Pi - values[0]))
}

