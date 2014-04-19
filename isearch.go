package main

import (
	"log"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
)

// Searchresult struct - holds for each search result
// title of programme, thumbnail url, and the programmes'
// iplayer index.
type Searchresult struct {
	Title         string
	Thumbnail     string
	Index         string
	Oldrecordings string
	OldRec        bool
}

// Category struct - holds name of an iplayer category
// plus the get-request url, e.g. ?category=<name>
type Category struct {
	Name string
	Url  string
}

func main() {
	RunServer()
}

// refresh get_iplayer's index
func init() {
	exec.Command("get_iplayer", "--refresh").Start()
}

// NewSearch - takes a map that contains either the category to search for, e.g. films,
// or a search string to run the get_iplayer command with.
// for every found match, a Search result struct gets initialized with the
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
	if isoldRec(string(isoOut)) {
		recstring := strings.Split(listOldRecordings(isoOut), "\n")
		for _, i := range recstring {
			if i == "Do you wish to delete them now (Yes/No) ?" {
				break
			}
			new := Searchresult{Oldrecordings: i}
			result = append(result, new)
		}

		return result
	}

	for _, i := range isoOutslice {
		new := Searchresult{Title: title(i), Index: index(i), Thumbnail: thumbnail(i)}
		if new != (Searchresult{}) {
			result = append(result, new)
		}
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

// index - takes an iplayer search result
// and returns the index consisting of a string
// of 1 - 4 digits.
// index("233 http://...jpg Some programme") -> "233"
func index(s string) string {
	re := regexp.MustCompile(`[0-9]*`)
	nonempty := ""
	if re.FindString(s) != "" {
		nonempty = "/info?index=" + re.FindString(s)
	}
	return nonempty
}

// title - takes an iplayer search result
// and returns the programme's title
// title("232 http://...jpg Some Programme") -> "Some Programme"
func title(s string) string {
	re := regexp.MustCompile("jpg [A-Z0-9].*")
	prelim := re.FindString(s)
	re = regexp.MustCompile("[A-Z0-9].*")
	return re.FindString(prelim)
}

// thumbnail - takes an iplayer search result
// and returns the url for it's thumbnail
// thumbnail("232 http://...jpg Some Programme") -> "http://...jpg"
func thumbnail(s string) string {
	re := regexp.MustCompile("http.*jpg")
	return re.FindString(s)
}

// isoldRec - if with get_iplayer downloaded programmes
// are stored for > 30 days, get_iplayer prints a
// list of all to be deleted programmes.
// isoldRec checks for the existence of this message.
func isoldRec(s string) bool {
	re := regexp.MustCompile("These programmes.*")
	return re.MatchString(s)
}

// listOldRecordings - returns a string of
// all downloaded programmes, that have been
// stored for > 30 days.
func listOldRecordings(s string) string {
	re := regexp.MustCompile("These programmes.*")
	pos := re.FindStringIndex(s)
	return s[pos[0]:]
}

// IplayerInfo - for every iplayer programme,
// this struct holds the url for a thumbnail of size 4,
// the programmes title, the long description and
// the available stream quality, e.g. flashhd
type IplayerInfo struct {
	Thumbnail, Description, Title, DownloadUrl, ImdbUrl string
	Modes                                               []string
	Pagetitle                                           string
	OldRecording                                        string
}

// IplayerIndex - every iplayer programme has a
// unique index (1 to 4 digits), to facilitate
// download and info about programmes.
type IplayerIndex struct {
	index string
}

// Thumb4 - find thumbnail of size 4 in iplayer search
// output.
func (i *IplayerIndex) Thumb4() string {
	re := regexp.MustCompile("thumbnail4.*")
	prelim := re.FindString(i.index)
	re = regexp.MustCompile("htt.*")
	return re.FindString(prelim)
}

// Description - find long description in iplayer search
// output.
func (i *IplayerIndex) Description() string {
	re := regexp.MustCompile("desc:.*")
	prelim := re.FindString(i.index)
	re = regexp.MustCompile("[A-Z].*")
	return re.FindString(prelim)
}

// Title - find title in iplayer search output.
func (i *IplayerIndex) Title() string {
	re := regexp.MustCompile("title:.*")
	prelim := re.FindString(i.index)
	re = regexp.MustCompile("[A-Z0-9].*")
	return re.FindString(prelim)
}

// Imdb - find nameshort in iplayer search output
// and add it to the url of imdb. Replace spaces with + sign
// to generate url-query url.
func (i *IplayerIndex) Imdb() string {
	re := regexp.MustCompile("nameshort:.*")
	prelim := re.FindString(i.index)
	re = regexp.MustCompile("[A-Z0-9].*")
	query := strings.Replace(re.FindString(prelim), " ", "+", -1)
	return "http://imdb.com/find?q=" + query
}

// Modes - Collect available Stream Download modes
// ranging from hd to low bitrate quality.
func (i *IplayerIndex) Modes() []string {
	re := regexp.MustCompile("modes.*")
	prelim := re.FindString(i.index)
	high := regexp.MustCompile("flashhigh")
	vhigh := regexp.MustCompile("flashvhigh")
	hd := regexp.MustCompile("flashhd")
	low := regexp.MustCompile("flashlow")
	modes := []string{high.FindString(prelim),
		vhigh.FindString(prelim), hd.FindString(prelim),
		low.FindString(prelim)}
	for _, i := range modes {
		if i == "" {
			strings.TrimSuffix(i, "")
		}
	}
	return modes
}

// IplayerInfoOutput - takes a string index digit(s)
// returns IplayerIndex struct with the output
// of the iplayer info command as a string.
func IplayerInfoOutput(s string) *IplayerIndex {
	info := exec.Command("get_iplayer", "-i", s)
	infoOut, err := info.Output()
	if err != nil {
		panic(err)
	}
	return &IplayerIndex{string(infoOut)}
}

// DownloadProgramme - takes an iplayer index
// and starts the download of said programmme by opening
// a gnome-terminal and invoking get_iplayer
func DownloadProgramme(index, mode string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("gnome-terminal", "--working-directory="+usr.HomeDir+"/Videos",
		"-e",
		"get_iplayer --modes="+mode+"1"+" -g "+index+" --flvstreamer=/usr/bin/flvstreamer")
	cmd.Start()
}
