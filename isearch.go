package main

import (
        "fmt"
        "os/exec"
        "regexp"
        "strings"
)

// Ipsearch struct - contains the get_iplayer output
// for a name-search, e.g. get_iplayer "silk"
type Ipsearch struct {
        searchterm string
}

// Catsearch struct: - contains the get_iplayer output
// for a category search, e.g. get_iplayer --category "films"
type Catsearch struct {
        Ipsearch
}

// Searchresult struct - holds for each searchresult
// title of programme, thumbnail url, and the programmes'
// iplayer index.
type Searchresult struct {
        title     string
        thumbnail string
        index     string
}

func main() {
        newsearch := newSearch("pramface")
        fmt.Println(newsearch.searchterm)
        cats := newCategory("legal")
        fmt.Println(cats.searchterm)
        fmt.Println(cats.ThumbNail())
        fmt.Println(cats.Index())
        cats.Title()
        fmt.Println(newsearch.Title())

}
func newSearch(s string) *Ipsearch {
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "\"<index> <thumbnail> <name> <episode>\"",
                s)
        isoOut, err := search.Output()
        if err != nil {
                panic(err)
        }
        inf := regexp.MustCompile(`INFO`)
        infpos := inf.FindStringIndex(string(isoOut))
        return &Ipsearch{searchterm: string(isoOut)[:(infpos[0])]}
}
func newCategory(cat string) *Catsearch {
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "\"<index> <thumbnail> <name> <episode>\"",
                "--category", cat)
        catOut, err := search.Output()
        if err != nil {
                panic(err)
        }
        inf := regexp.MustCompile(`INFO`)
        infpos := inf.FindStringIndex(string(catOut))
        return &Catsearch{Ipsearch{searchterm: string(catOut)[:(infpos[0])]}}
}

// ThumbNail - return string of thumbnail url in search result.
func (ip *Ipsearch) ThumbNail() []string {
        re := regexp.MustCompile("http.*jpg")
        return re.FindAllString(ip.searchterm, -1)
}

// Index - return string of the index in search result in
// form of digits.
func (ip *Ipsearch) Index() []string {
        re := regexp.MustCompile(`"[0-9]*`)
        s := re.FindAllString(ip.searchterm, -1)
        result := make([]string, 0)
        for _, i := range s {
                i = strings.TrimSpace(i[1:])
                result = append(result, i)
        }
        return removeEmpty(result)
}

// Title return string of the programmes title
// in the search result.
func (ip *Ipsearch) Title() []string {
        re := regexp.MustCompile(`(\s[A-Z0-9].[^"]*)`)
        s := re.FindAllString(ip.searchterm, -1)
        return removeEmpty(s)
}
func removeEmpty(s []string) []string {
        empty := make([]string, 0)
        for _, i := range s {
                if i != "" {
                        empty = append(empty, i)
                }
        }
        return empty
}
