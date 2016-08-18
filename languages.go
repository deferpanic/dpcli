package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

// FIXME
type Language struct {
	Name string
}

type LanguagesResponse struct {
	Title     string
	Error     string
	Languages []Language
}

type Languages struct{}

func (languages *Languages) List() {
	lr := LanguagesResponse{}
	err := cli.GetJSON(languagesURL, &lr)
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(lr.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("Name")})

		// FIXME - auto-format
		for i := 0; i < len(lr.Languages); i++ {
			table.Append([]string{
				lr.Languages[i].Name})
		}

		table.Render()

	}

}
