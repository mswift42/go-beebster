package main

import (
        "github.com/codegangsta/martini"
        "github.com/martini-contrib/render"
)

func RunServer() {
        cats := Categoryinit()
        m := martini.Classic()
        m.Use(render.Renderer(render.Options{Layout: "layout",
                Directory: "templates"}))
        m.Get("/about", func(r render.Render) {
                r.HTML(200, "about", cats)

        })
        // m.Get("/about", func(r render.Render) {
        //         return "Hallo"
        // })

        m.Run()
}
