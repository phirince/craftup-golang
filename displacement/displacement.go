package main

import (
  "fmt"
  "math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
  return func(t float64) float64 {
    return 0.5 * a * math.Pow(t, 2) + v0 * t + s0
  }
}

func main() {
  func1 := GenDisplaceFn(2, 10, 100)
  func2 := GenDisplaceFn(3, 5, 20)
  fmt.Println(func1(5))
  fmt.Println(func2(5))
}
