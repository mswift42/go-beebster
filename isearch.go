package main

import (
        "fmt"
        "os/exec"
        "regexp"
        "strings"
)

// Searchresult struct - holds for each searchresult
// title of programme, thumbnail url, and the programmes'
// iplayer index.
type Searchresult struct {
        title         string
        thumbnail     string
        index         string
        oldrecordings string
}

// Category struct - holds name of an iplayer category
// plus the get-request url, e.g. ?category=<name>
type Category struct {
        name string
        url  string
}

func main() {
        newsearch := NewSearch(map[string]string{"category": "", "searchvalue": "pramface"})
        cats := NewSearch(map[string]string{"category": "legal"})
        newcats := NewSearch(map[string]string{"category": "films"})
        fmt.Println(newsearch)
        fmt.Println(cats)
        fmt.Println(newcats)
        fmt.Println(NewCategory("films"))
        //        RunServer()
}

// NewSearch - takes a map that contains either the category to search for, e.g. films,
// or a searchstring to run the get_iplayer command with.
// for every found match, a Searchresult struct gets initialized with the
// index, thumbnail and title for the given match.
func NewSearch(s map[string]string) []Searchresult {
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

// searchResults - helper fn for NewSearch.
// if given map contains a category, run the iplayer cmd
// with --category <category>
// else invoke iplayer search with s["category":"","searchvalue":<searchterm>
func searchResult(s map[string]string) (string, error) {
        if s["category"] == "" {
                search := exec.Command("get_iplayer", "--nocopyright",
                        "--limitmatches", "50",
                        "--listformat", "<index> <thumbnail> <name> <episode>",
                        s["searchvalue"])
                isoOut, err := search.Output()
                return string(isoOut), err
        }
        search := exec.Command("get_iplayer", "--nocopyright",
                "--limitmatches", "50",
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

// Categories - Array with iplayer categories.
var Categories = []string{"popular", "highlights", "films", "nature", "crime",
        "sitcom", "thriller", "legal", "sport"}

// NewCategory - takes a description of a iplayer category.
// Returns a new instance of struct category with the name and
// the url for a category, e.g. *Category{name:"films",url:"?category=films"}
func NewCategory(s string) *Category {
        cat := Category{name: s,
                url: ("?category=" + s)}
        return &cat
}
