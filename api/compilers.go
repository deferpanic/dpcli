package api

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

type Compiler struct {
	Name string
}

type CompilersResponse struct {
	Title     string
	Error     string
	Compilers []Compiler
}

type Compilers struct{}

func (compilers *Compilers) List() {
	cr := CompilersResponse{}

	err := Cli.GetJSON(APIBase+"/compilers", &cr)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
	} else {
		fmt.Println(GreenBold(cr.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{GreenBold("Name")})

		// FIXME - auto-format
		for i := 0; i < len(cr.Compilers); i++ {
			table.Append([]string{cr.Compilers[i].Name})
		}

		table.Render()

	}

}
