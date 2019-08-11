package main

// import "fmt"

import (
  "syscall/js"
)

func add(i []js.Value) {
  js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
  println(js.ValueOf(i[0].Int() + i[1].Int()).String())
}

func subtract(i []js.Value) {
  js.Global().Set("output", js.ValueOf(i[0].Int()-i[1].Int()))
  println(js.ValueOf(i[0].Int() - i[1].Int()).String())
}

func registerCallbacks() {
  js.Global().Set("add", js.NewCallback(add))
  js.Global().Set("subtract", js.NewCallback(subtract))
}

func main() {
  c := make(chan struct{}, 0)

  println("Go WebAssembly Initialized :)")
  registerCallbacks()
  <-c
}
