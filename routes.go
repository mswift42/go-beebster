package main

import (
        "github.com/codegangsta/martini"
        "github.com/martini-contrib/render"
        "net/http"
)

type result struct {
        searchvalue string `form:"searchvalue"`
}

func RunServer() {
        //        cats := Categoryinit()
        m := martini.Classic()
        m.Use(render.Renderer(render.Options{Layout: "layout",
                Directory: "templates"}))
        m.Get("/", func(r render.Render) {
                r.HTML(200, "index", "index")
        })
        m.Get("/about", func(r render.Render) {
                r.HTML(200, "about", "about")

        })

        m.Post("/results", func(r render.Render, re *http.Request) {
                sv := re.FormValue("searchvalue")
                resmap := map[string]string{"category": "",
                        "searchvalue": sv}
                out := NewSearch(resmap)
                r.HTML(200, "result", out)
        })
        m.Get("/info", func(r render.Render, re *http.Request) {
                info := re.URL.Query().Get("index")

                ind := IplayerInfoOutput(info)
                iplayerinfo := &IplayerInfo{Thumbnail: ind.Thumb4(),
                        Description: ind.Description(),
                        Title:       ind.Title(),
                        DownloadUrl: "/download?index=" + info}
                r.HTML(200, "info", iplayerinfo)
        })
        m.Run()
}
