package main

import (
  "fmt"
)

func Swap(sli []int, index int) {
  sli[index], sli[index+1] = sli[index+1], sli[index]
}

func BubbleSort (sli []int) {
  for i := 0; i < len(sli) - 1; i++ {
    for j := 0; j < len(sli) - i - 1; j++ {
      if sli[j] > sli[j+1] {
        Swap(sli, j)
      }
    }
  }
}

func main() {
  var a = []int{1, 3, 2, 5, 4, 100, 323, 44, 0, 12}
  fmt.Println(a)
  BubbleSort(a)
  fmt.Println(a)
  var b = []int{1, 9, -9, 24, 4, -1, 323, 44, 0, 12}
  BubbleSort(b)
  fmt.Println(b)
}

