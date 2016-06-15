package main

import (
	"fmt"
)

type Builtins struct{}

func (builtins *Builtins) List() {
	response, err := cli.Postit(nil, builtinsURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}
