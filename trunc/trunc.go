package main
import "fmt"

func main () {
  var input float32;
  var dec int;
  var num, err = fmt.Scan(&input, 1);
  if num < 1 {
    fmt.Printf("No input read %s\n", err);
    return;
  }
  dec = int(input);
  fmt.Printf("Truncated value: %d\n", dec)
}
