package main

import (
        "fmt"
        "os/exec"
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

var iplayerSearchCommand = "get_iplayer --nocopyright --limitmatches 50 --listformat \"<index> <pid> <thumbnail> <name> <episode>\""

func main() {
        ipCmd := exec.Command("get_iplayer", "--category", "films")
        ipOut, err := ipCmd.Output()
        if err != nil {
                panic(err)
        }
        fmt.Println(string(ipOut))
        newsearch := newSearch("silk")
        fmt.Println(newsearch.searchterm)
        cats := newCategory("legal")
        fmt.Println(cats.searchterm)

}
func newSearch(s string) *Ipsearch {
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "\"<index> <pid> <thumbnail> <name> <episode>\"",
                s)
        isoOut, err := search.Output()
        if err != nil {
                panic(err)
        }
        return &Ipsearch{searchterm: string(isoOut)}
}
func newCategory(cat string) *Catsearch {
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "\"<index> <pid> <thumbnail> <name> <episode>\"",
                "--category", cat)
        catOut, err := search.Output()
        if err != nil {
                panic(err)
        }
        return &Catsearch{Ipsearch{searchterm: string(catOut)}}
}
