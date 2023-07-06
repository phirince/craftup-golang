package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
  var file_name string
  type Name struct {
    fname string
    lname string
  }
  name_list := []Name{}
  fmt.Print("Enter filename: ")
  fmt.Scan(&file_name)
  f_handle, err := os.Open(file_name)
  if err != nil {
    fmt.Printf("Error opening file: %s\n", err)
  }
  scanner := bufio.NewScanner(f_handle)
  for scanner.Scan() {
    res := strings.Split(scanner.Text(), " ")
    name_list = append(name_list, Name{fname: res[0], lname: res[1]})
  }
  f_handle.Close()
  for _, name := range name_list {
    fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
  }
}
