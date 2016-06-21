package main

import (
	"fmt"
)

type Addons struct{}

func (addons *Addons) List() {
	response, err := cli.Postit(nil, addonsURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}

func (addons *Addons) New() {
}

func (addons *Addons) Delete() {
}
