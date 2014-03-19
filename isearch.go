package main

import (
        "fmt"
        "os/exec"
        "regexp"
        "strings"
)

// // Ipsearch struct - contains the get_iplayer output
// // for a name-search, e.g. get_iplayer "silk"
// type Ipsearch struct {
//         searchterm []string
// }

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
        cats := newSearch(map[string]string{"category": "legal"})
        newcats := newSearch(map[string]string{"category": "films"})
        fmt.Println(newsearch)
        fmt.Println(cats)
        fmt.Println(newcats)
        fmt.Println(index("871 http://www.bbc.co.uk/iplayer/images/episode/b03ws0yr_150_84.jpg Pramface: Series 3 1. I'm Excited Too!"))
        fmt.Println(title("871 http://www.bbc.co.uk/iplayer/images/episode/b03ws0yr_150_84.jpg Pramface: Series 3 1. I'm Excited Too!"))
        fmt.Println(thumbnail("871 http://www.bbc.co.uk/iplayer/images/episode/b03ws0yr_150_84.jpg Pramface: Series 3 1. I'm Excited Too!"))

}
func newSearch(s map[string]string) []Searchresult {
        isoOut, err := searchResult(s)
        if err != nil {
                panic(err)
        }
        inf := regexp.MustCompile(`INFO`)
        infpos := inf.FindStringIndex(string(isoOut))
        isoOutslice := strings.Split(
                strings.Replace(string(isoOut)[:infpos[0]], "Matches:", "", 1), "\n")
        result := make([]Searchresult, 0)
        for _, i := range isoOutslice {
                new := Searchresult{title: title(i), index: index(i), thumbnail: thumbnail(i)}
                result = append(result, new)
        }
        return result
}
func searchResult(s map[string]string) (string, error) {
        if s["category"] == "" {
                search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                        "--listformat", "<index> <thumbnail> <name> <episode>",
                        s["searchvalue"])
                isoOut, err := search.Output()
                return string(isoOut), err
        }
        search := exec.Command("get_iplayer", "--nocopyright", "--limitmatches", "50",
                "--listformat", "<index> <thumbnail> <name> <episode>",
                "--category", s["category"])
        isoOut, err := search.Output()
        return string(isoOut), err
}

func index(s string) string {
        re := regexp.MustCompile(`[0-9]*`)
        return re.FindString(s)
}
func title(s string) string {
        re := regexp.MustCompile("jpg [A-Z0-9].*")
        prelim := re.FindString(s)
        re = regexp.MustCompile("[A-Z0-9].*")
        return re.FindString(prelim)
}
func thumbnail(s string) string {
        re := regexp.MustCompile("http.*jpg")
        return re.FindString(s)
}
