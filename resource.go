package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Resource is the base struct for management of resources
type Resource struct {
	Name    string `json:"Name"`
	Owner   string `json:"Owner"`
	Builtin string `json:"Builtin"`
}

type Resources struct{}

func (resources *Resources) New(name string, owner string, builtin string) {
	resource := Resource{}
	resource.Name = name
	resource.Owner = owner
	resource.Builtin = builtin

	b, err := json.Marshal(resource)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, newresourceURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}

func (resources *Resources) List(name string) {
	resource := Resource{}
	resource.Owner = name

	b, err := json.Marshal(resource)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(b, listresourcesURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}
