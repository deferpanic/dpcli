package api

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

// New provisions and attaches a resource to a project
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

	response, err := Cli.Postit(b, APIBase+"/resource/new")
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}

}

// Available lists the resources available
func (resources *Resources) Available() {
	response, err := Cli.Postit(nil, systemURL+"/resources")
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}
}

// List lists the resources provisioned
func (resources *Resources) List() {
	resource := Resource{}

	b, err := json.Marshal(resource)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := Cli.Postit(b, APIBase+"/resource/list")
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}

}

// ListByName lists the resources provisioned to project_name
func (resources *Resources) ListByName(name string) {
	resource := Resource{}
	resource.Owner = name

	b, err := json.Marshal(resource)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := Cli.Postit(b, APIBase+"/resource/list")
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}

}
