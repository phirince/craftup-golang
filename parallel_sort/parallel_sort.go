package main

import "fmt"
import "sort"

func sort_int(arr []int, c chan []int) {
  sort.Ints(arr)
  c <- arr
}

func merge_arrays(arr1 []int, arr2 []int) []int {
  sorted_arr := []int{}
  for _, v := range arr1{
    j := 0
    for ; j < len(arr2); j++ {
      w := arr2[j]
      if w <= v {
        sorted_arr = append(sorted_arr, w)
      } else {
        break
      }
    }
    arr2 = arr2[j:]
    sorted_arr = append(sorted_arr, v)
  }
  sorted_arr = append(sorted_arr, arr2...)
  return sorted_arr
}
func main() {
  arr := []int{5, 10, 24, 100, -1, 12, 45, 50, -200, 10023423, 9, 15, 45}
  length := len(arr)
  fmt.Println("Length of slice", length)
  c := make(chan []int)
  go sort_int(arr[0:(length/4)], c)
  go sort_int(arr[(length/4):(length/2)], c)
  go sort_int(arr[(length/2):(length*3/4)], c)
  go sort_int(arr[(length*3/4):], c)
  arr1 := <- c
  arr2 := <- c
  sorted_arr := merge_arrays(arr1, arr2)
  arr3 := <- c
  sorted_arr = merge_arrays(sorted_arr, arr3)
  arr4 := <- c
  sorted_arr = merge_arrays(sorted_arr, arr4)
  fmt.Println(sorted_arr)
}
