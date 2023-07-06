package main
import "fmt"
import "bufio"
import "os"
import "strings"

func main ()  {
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print("Enter the text: ")
  scanner.Scan()
  line := scanner.Text()
  // fmt.Printf("Input test: %s", line)
  if strings.HasPrefix(strings.ToLower(line), "i") && strings.HasSuffix(strings.ToLower(line), "n") && strings.Contains(strings.ToLower(line), "a") {
    fmt.Print("Found!\n")
  } else {
    fmt.Print("Not Found!\n")
  }
}
