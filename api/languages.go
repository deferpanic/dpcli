package api

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

type Language struct {
	ID       int
	Name     string
	Version  string
	Compiler string
}

type LanguagesResponse struct {
	Title     string
	Error     string
	Languages []Language
}

type Languages struct{}

func (languages *Languages) List() {
	lr := LanguagesResponse{}

	err := Cli.GetJSON(APIBase+"/languages", &lr)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
	} else {
		fmt.Println(GreenBold(lr.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{GreenBold("Name"),
			GreenBold("Version"),
			GreenBold("Compiler")})

		// FIXME - auto-format
		for i := 0; i < len(lr.Languages); i++ {
			table.Append([]string{lr.Languages[i].Name,
				lr.Languages[i].Version,
				lr.Languages[i].Compiler})
		}

		table.Render()

	}

}
