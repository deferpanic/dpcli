package main

import ()

type Status struct{}

func (status *Status) Show() {
	response, err := cli.Postit(nil, statusURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}
