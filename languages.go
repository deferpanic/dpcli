package main

import ()

type Languages struct{}

func (languages *Languages) List() {
	response, err := cli.Postit(nil, languagesURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}
