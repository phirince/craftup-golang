package main

import "fmt"

type Animal interface {
  Eat()
  Move()
  Speak()
}

type Cow struct {
  name string
}

func (c *Cow) Eat() {
  fmt.Println("grass")
}

func (c *Cow) Move() {
  fmt.Println("walk")
}

func (c *Cow) Speak() {
  fmt.Println("moo")
}

type Bird struct {
  name string
}

func (c *Bird) Eat() {
  fmt.Println("worms")
}

func (c *Bird) Move() {
  fmt.Println("fly")
}

func (c *Bird) Speak() {
  fmt.Println("peep")
}

type Snake struct {
  name string
}

func (c *Snake) Eat() {
  fmt.Println("mice")
}

func (c *Snake) Move() {
  fmt.Println("slither")
}

func (c *Snake) Speak() {
  fmt.Println("hsss")
}

func main()  {
  var command string
  var arg1 string
  var arg2 string
  animal_list := []Animal{}
  for {
    fmt.Print("> ")
    num, _ := fmt.Scan(&command, &arg1, &arg2)
    if command == "newanimal" {
      if num < 3 {
        fmt.Println("newanimal command requires 3 arguments")
        continue
      }
      switch arg2 {
      case "cow": animal_list = append(animal_list, &Cow{name: arg1})
      case "bird": animal_list = append(animal_list, &Bird{name: arg1})
      case "snake": animal_list = append(animal_list, &Snake{name: arg1})
      default: fmt.Println("Unrecognized animal:", arg2); continue 
      }
      fmt.Println("Created it!")
    } else if command == "query" {
      if num < 3 {
        fmt.Println("query command requires 3 arguments")
        continue
      }
      var animal Animal
      var c *Cow
      var b *Bird
      var s *Snake
      switch arg1 {
      case "cow": animal = c
      case "bird": animal = b
      case "snake": animal = s
      default: fmt.Println("Unrecognized animal:", arg2); continue 
      }
      switch arg2 {
      case "eat": animal.Eat()
      case "move": animal.Move()
      case "speak": animal.Speak()
      default: fmt.Println("Unrecognized animal:", arg2); continue 
      }
    } else {
      fmt.Println("Invalid command:", command)
    }
  }
}
