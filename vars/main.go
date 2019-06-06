package main

import "fmt"

func main() {
    var age int32 = 27
    const isCool bool = true

    // Shorthand
    name, height := "Mac", 1.76
    fmt.Println(name, age, isCool, height)
}
