package main

import (
        "fmt"
        "os/exec"
        "regexp"
)

type Ipsearch struct {
        searchterm    string
        title         string
        thumbnail     string
        index         string
        oldrecordings bool
}
type Catsearch struct {
        Ipsearch
}

func main() {
        newsearch := newSearch("silk")
        fmt.Println(newsearch.searchterm)
        cats := newCategory("legal")
        fmt.Println(cats.searchterm)
        fmt.Println(cats.ThumbNail())

}
func newSearch(s string) *Ipsearch {
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "\"<index> <thumbnail> <name> <episode>\"",
                s)
        isoOut, err := search.Output()
        if err != nil {
                panic(err)
        }
        return &Ipsearch{searchterm: string(isoOut)}
}
func newCategory(cat string) *Catsearch {
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "\"<index> <thumbnail> <name> <episode>\"",
                "--category", cat)
        catOut, err := search.Output()
        if err != nil {
                panic(err)
        }
        return &Catsearch{Ipsearch{searchterm: string(catOut)}}
}

func (ip *Ipsearch) ThumbNail() []string {
        re := regexp.MustCompile("http.*jpg")
        return re.FindAllString(ip.searchterm, -1)
}
