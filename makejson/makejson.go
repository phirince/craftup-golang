package main
import (
  "os"
  "bufio"
  "fmt"
  "encoding/json"
)

func main () {
  person := make(map[string]string)
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Enter name: ")
  scanner.Scan()
  person["name"] = scanner.Text()
  fmt.Print("Enter address: ")
  scanner.Scan()
  person["address"] = scanner.Text()

  json_string, err := json.Marshal(person) 
  if err != nil {
    fmt.Printf("Error marshalling data: %v", err)
  }

  fmt.Printf("%s\n", json_string)
}
