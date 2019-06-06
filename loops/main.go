package main

import "fmt"

func main() {
  i := 1

  for i <= 10 {
    fmt.Println(i)
    i++
  }

  for e := 1; e <= 10; e++ {
    fmt.Printf("Number %d\n", e)
  }

  // FizzBuzz
  for q := 1; q <= 100; q++ {
    if q % 15 == 0 {
      fmt.Println("FizzBuzz")
    } else if q % 3 == 0 {
      fmt.Println("Fizz")
    } else if q % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(q)
    }
  }
}
