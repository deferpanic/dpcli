package main

import (
	"encoding/json"
	"log"
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
		log.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(b, scaleupURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}

func (instances *Instances) Log(domain string) {
	instance := &Instance{}
	instance.Domain = domain
	b, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, runlogURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}

func (instances *Instances) List(name string) {
	image := &Image{}
	image.Name = name

	b, err := json.Marshal(image)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, showURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}

func (instances *Instances) Pause(domain string) {
	instance := &Instance{}
	instance.Domain = domain

	b, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, pauseURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}

func (instances *Instances) Resume(domain string) {
	instance := &Instance{}
	instance.Domain = domain

	b, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, resumeURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}

func (instances *Instances) Scaleup(name string, count int) {
	image := &Image{}
	image.Name = name
	if count < 1 {
		log.Println("Number of image instances to launch can't be less than 1")
		os.Exit(1)
	}
	image.Count = count

	// FIXME
	image.Force = true

	b, err := json.Marshal(image)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, scaleupURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}

func (instances *Instances) Scaledown(name string, domain string) {
	instance := &Instance{}
	instance.Name = name
	instance.Domain = domain

	// FIXME
	instance.Force = true

	b, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, scaledownURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}
