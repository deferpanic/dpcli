package main

import (
	"fmt"
)

type Languages struct{}

func (languages *Languages) List() {
	response, err := cli.Postit(nil, languagesURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}
