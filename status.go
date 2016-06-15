package main

import (
	"fmt"
)

type Status struct{}

func (status *Status) Show() {
	response, err := cli.Postit(nil, statusURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}
