package api

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

type Search struct{}

type SearchItem struct {
	Name        string
	Description string
	Stars       int
	Language    string
	Source      string
	CreatedAt   time.Time
}

type SearchResponse struct {
	Title       string
	Error       string
	SearchItems []SearchItem
}

// Find lists public projects w/a name or a description that match description
func (search *Search) Find(description string) {
	sr := SearchResponse{}
	err := Cli.GetJSON(APIBase+"/projects/search/"+description, &sr)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
	} else {
		search.Format(sr)
	}
}

// Find lists public projects w/a name or a description that match description
// && have at least star cnt
func (search *Search) FindWithStars(description string, starcnt int) {
	s := strconv.Itoa(starcnt)

	sr := SearchResponse{}
	err := Cli.GetJSON(APIBase+"/projects/search/"+description+"/by_stars/"+s, &sr)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
	} else {
		search.Format(sr)
	}
}

func (search *Search) Format(sr SearchResponse) {
	fmt.Println(GreenBold(sr.Title))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)

	table.SetHeader([]string{GreenBold("Name"),
		GreenBold("Description"), GreenBold("Stars"), GreenBold("Language"),
		GreenBold("Source")})

	for i := 0; i < len(sr.SearchItems); i++ {
		istars := strconv.Itoa(sr.SearchItems[i].Stars)

		table.Append([]string{
			sr.SearchItems[i].Name,
			sr.SearchItems[i].Description,
			istars,
			sr.SearchItems[i].Language,
			sr.SearchItems[i].Source,
		})
	}

	table.Render()
}
