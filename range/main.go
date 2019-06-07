package main

import "fmt"

func main() {
  ids := []int{33,76,54,23,11,2}

  for i, id := range ids {
    fmt.Printf("%d - ID: %d\n", i, id)
  }

  for _, id := range ids {
    fmt.Printf("%d\n", id)
  }

  sum := 0

  for _, id := range ids {
    sum += id
  }

  fmt.Println("Sum", sum)

  emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

  for k, v := range emails {
    fmt.Printf("%s: %s\n", k, v)
  }
}
