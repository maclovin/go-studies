package main

import "fmt"

func main() {
  fruitArr := [3]string{"Apple", "Orange", "Grape"}

  fmt.Println(fruitArr, len(fruitArr), fruitArr[1:2])
}
