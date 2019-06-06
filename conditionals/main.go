package main

import "fmt"

func main() {
  x := 11
  y := 10

  if x < y {
    fmt.Printf("%d is less than %d\n", x, y)
  } else {
    fmt.Printf("%d is less than %d\n", y, x)
  }

  color := "black"

  if color == "red" {
    fmt.Println("color is red")
  } else if color == "blue" {
    fmt.Println("color is blue")
  } else {
    fmt.Println("color is not blue or red")
  }

  switch color {
    case "red":
      fmt.Println("color is red")
    case "blue":
      fmt.Println("color is blue")
    default:
      fmt.Println("color is not blue or red")
  }
}
