package main

import (
        "github.com/codegangsta/martini"
        "github.com/martini-contrib/render"
)

func RunServer() {
        m := martini.Classic()
        m.Get("/", func() string {
                return "Hello World!"
        })
        m.Run()
}
