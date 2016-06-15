package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Instance is the base struct for management of instances
type Instance struct {
	Name   string `json:"Name"`
	Domain string `json:"Domain"`
	Force  bool   `json:"Force"`
}

type Instances struct {
}

func (instances *Instances) New(name string) {
	image := &Image{}
	image.Name = name
	image.Count = 1

	// FIXME
	image.Force = true

	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(b, scaleupURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}

func (instances *Instances) Log(domain string) {
	instance := &Instance{}
	instance.Domain = domain
	b, err := json.Marshal(instance)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, runlogURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}

func (instances *Instances) List(name string) {
	image := &Image{}
	image.Name = name

	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, showURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}

func (instances *Instances) Pause(domain string) {
	instance := &Instance{}
	instance.Domain = domain

	b, err := json.Marshal(instance)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, pauseURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}

func (instances *Instances) Resume(domain string) {
	instance := &Instance{}
	instance.Domain = domain

	b, err := json.Marshal(instance)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, resumeURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}

// FIXME - wrong api
func (instances *Instances) ScaleUp(name string) {
	image := &Image{}
	image.Name = name
	image.Count = 1

	// FIXME
	image.Force = true

	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, scaleupURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}

// FIXME - wrong api
func (instances *Instances) ScaleDown(name string, domain string) {
	instance := &Instance{}
	instance.Name = name
	instance.Domain = domain

	// FIXME
	instance.Force = true

	b, err := json.Marshal(instance)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, scaledownURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}
