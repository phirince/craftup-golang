package main

import "fmt"
import "strconv"

func main() {
  inputs := make([]int, 0, 3)
  for {
    fmt.Print("Input Integer: ")
    var input_str string 
    _, err := fmt.Scan(&input_str)
    if err != nil {
      fmt.Printf("Error scanning input: %b", err)
      break
    }
    if input_str == "X" {
      break
    }
    input_int, err := strconv.Atoi(input_str)
    if err != nil {
      fmt.Printf("Invalid input: %s", input_str)
      break
    }
    inputs = append(inputs, input_int)
    for i := len(inputs) - 1; i > 0; i-- {
      if inputs[i] < inputs[i-1] {
        tmp := inputs[i-1]
        inputs[i-1] = inputs[i]
        inputs[i] = tmp
      }
    }
    fmt.Printf("%v", inputs)
  }
}
