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
        searchterm []string
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
        newsearch := newSearch(map[string]string{"category": "", "searchvalue": "pramface"})
        fmt.Println(newsearch.searchterm)
        cats := newSearch(map[string]string{"category": "sitcom"})
        fmt.Println(cats.searchterm)
        fmt.Println(cats.ThumbNail())
        fmt.Println(cats.Index())
        cats.Title()
        fmt.Println(newsearch.Title())
        newcats := newSearch(map[string]string{"category": "films"})
        fmt.Println(newcats.Title())

}
func newSearch(s map[string]string) *Ipsearch {
        isoOut, err := searchResult(s)
        if err != nil {
                panic(err)
        }
        inf := regexp.MustCompile(`INFO`)
        infpos := inf.FindStringIndex(string(isoOut))
        isoOutslice := strings.Split(
                strings.Replace(string(isoOut)[:infpos[0]], "Matches:", "", 1), "\n")
        return &Ipsearch{searchterm: isoOutslice}
}
func searchResult(s map[string]string) (string, error) {
        if s["category"] == "" {
                search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                        "--listformat", "\"<index> <thumbnail> <name> <episode>\"",
                        s["searchvalue"])
                isoOut, err := search.Output()
                return string(isoOut), err
        }
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "\"<index> <thumbnail> <name> <episode>\"",
                "--category", s["category"])
        isoOut, err := search.Output()
        return string(isoOut), err
}

func applySearch(s []string, pat string) []string {
        re := regexp.MustCompile(pat)
        result := make([]string, 0)
        for _, i := range s {
                result = append(result, re.FindString(i))
        }
        return result
}

// ThumbNail - return string of thumbnail url in search result.
func (ip *Ipsearch) ThumbNail() []string {
        pat := "http.*jpg"
        return (applySearch(ip.searchterm, pat))
}

// Index - return string of the index in search result in
// form of digits.
func (ip *Ipsearch) Index() []string {
        pat := (`"[0-9]*`)
        slice := applySearch(ip.searchterm, pat)
        result := make([]string, 0)
        for _, i := range slice {
                result = append(result, (strings.Replace(i, "\"", "", -1)))

        }

        return result
}

// Title return string of the programmes title
// in the search result.
func (ip *Ipsearch) Title() []string {
        pat := (`(\s[A-Z0-9].[^"]*)`)
        return applySearch(ip.searchterm, pat)
}
