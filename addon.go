package main

import ()

type Addons struct{}

func (addons *Addons) List() {
	response, err := cli.Postit(nil, addonsURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}
