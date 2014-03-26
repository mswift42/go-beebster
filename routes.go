package main

import (
        "github.com/codegangsta/martini"
        "github.com/martini-contrib/render"
        "net/http"
)

func RunServer() {
        //        cats := Categoryinit()
        m := martini.Classic()
        m.Use(render.Renderer(render.Options{Layout: "layout",
                Directory: "templates"}))
        m.Get("/", func(r render.Render) {
                r.HTML(200, "index", map[string]interface{}{"Pagetitle": "Search"})
        })
        m.Get("/about", func(r render.Render) {
                r.HTML(200, "about",
                        map[string]interface{}{"Pagetitle": "About"})

        })

        m.Post("/results", func(r render.Render, re *http.Request) {
                sv := re.FormValue("searchvalue")
                resmap := map[string]string{"category": "",
                        "searchvalue": sv}
                out := NewSearch(resmap)
                r.HTML(200, "result",
                        map[string]interface{}{"Pagetitle": "Results",
                                "results": out})
        })
        m.Get("/info", func(r render.Render, re *http.Request) {
                info := re.URL.Query().Get("index")

                ind := IplayerInfoOutput(info)
                iplayerinfo := &IplayerInfo{Thumbnail: ind.Thumb4(),
                        Description: ind.Description(),
                        Title:       ind.Title(),
                        Pagetitle:   "Info",
                        DownloadUrl: "/download?index=" + info,
                        Modes:       ind.Modes(),
                        ImdbUrl:     ind.Imdb()}
                r.HTML(200, "info", iplayerinfo)
        })
        m.Get("/categories", func(r render.Render, re *http.Request) {
                cat := re.URL.Query().Get("category")
                catmap := map[string]string{"category": cat}
                out := NewSearch(catmap)
                r.HTML(200, "result",
                        map[string]interface{}{"Pagetitle": "Categories",
                                "results": out})
        })
        m.Any("/download", func(r render.Render, re *http.Request) {
                index := re.URL.Query().Get("index")
                mode := re.FormValue("mode")
                r.HTML(200, "download", map[string]interface{}{"Pagetitle": "Download"})
                DownloadProgramme(index, mode)
        })
        m.Run()
}
