package api

import (
	"fmt"
)

type Status struct{}

func (status *Status) Show() {
	response, err := Cli.Postit(nil, statusURL)
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}
}
