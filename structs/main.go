package main

import (
  "fmt"
  "strconv"
)

type Person struct {
  firstName, lastName, city, gender string
  age int
}

func (p Person) greet() string {
  return "Hello, my name id " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
  p.age++
}

func (p *Person) getMarried(partnerLastName string) {
  if p.gender == "f" {
    return
  } else {
    p.lastName += " " + partnerLastName
  }
}

func main() {
  person1 := Person{"Murilo", "Mac", "São Paulo", "m", 27}

  fmt.Println(person1.firstName)
  person1.age++
  fmt.Println(person1.age)
  person1.hasBirthday()
  person1.getMarried("Lol")
  fmt.Println(person1.greet())
}
