package main

import "fmt"

func main() {
  emails := make(map[string]string)

  emails["MacLovin"] = "maclovin@contaquanto.com.br"
  emails["Murilo Mac"] = "m4k2005@gmail.com"
  emails["Morolo Moc"] = "bololo@bololo.lo"

  delete(emails, "Morolo Moc")
  fmt.Println(len(emails), emails["Murilo Mac"])

  skills := map[string]string{"MacLovin": "Developer", "Murilo Mac": "Music Production"}

  fmt.Println(skills)
}
