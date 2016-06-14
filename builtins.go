package main

import ()

type Builtins struct{}

func (builtins *Builtins) List() {
	response, err := cli.Postit(nil, builtinsURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}
