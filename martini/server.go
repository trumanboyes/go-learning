package main

import "github.com/codegangsta/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })
  m.Get("/tru", func() string {
   return "<h1> Tru Function </h1>"
  })
  m.Run()
}
