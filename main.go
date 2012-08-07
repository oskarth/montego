package main

import (
  "math"
  "math/rand"
  "time"
  "fmt"
)

// unit circle and unit square

func withinCircle(x float64, y float64) bool {
  return ((x * x + y * y) <= 1)
}

func main() {
  // seed the random number generator with the current time
  rand.Seed(time.Now().Unix())
  circleArea := 0.0
  totalPoints := 0.0

  for i := 1; i < 1000000; i++ {
    var x = rand.Float64()
    var y = rand.Float64()

    if (withinCircle(x,y)) {
      circleArea += 1.0
    }
    totalPoints += 1.0
  }

  var approxPi = 4.0 * (circleArea / totalPoints)
  var deltaP = math.Abs(math.Pi - approxPi)
  fmt.Print(approxPi, " ", deltaP, "\n")
}
