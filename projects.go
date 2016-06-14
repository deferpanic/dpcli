package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

type Projects struct{}

// New creates a new project
func (projects *Projects) New(name string, language string, script string) {
	image := &Image{}
	image.Name = name
	image.Language = language
	image.MakeBin = true

	data, err := ioutil.ReadFile(script)
	if err == nil {
		image.Script = string(data)
	}

	b, err := json.Marshal(image)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, newURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}
}

// NewFromImage creates a new project from an image
func (projects *Projects) NewFromImage(name string, imagePath string) {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(data, putURL+"/"+url.QueryEscape(name))
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}

// List lists all your projects
func (projects *Projects) List() {
	response, err := cli.Postit(nil, displayURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}

// Log shows the log output for your project
func (projects *Projects) Log(name string) {
	image := &Image{}

	image.Name = name
	b, err := json.Marshal(image)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(b, makelogURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}

// ScaleUp scales up a project
func (projects *Projects) ScaleUp() {}

// Scaledown scales down a project
func (projects *Projects) ScaleDown() {}

// Download downloads everything in a project
// {image, volumes, deferpanic.yml if present
func (projects *Projects) Download(name string) {

	image := &Image{}
	image.Name = name
	b, err := json.Marshal(image)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, getURL)
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}

// Upload uploads a project
func (projects *Projects) Upload(name string, binary string) {
	data, err := ioutil.ReadFile(binary)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(data, putURL+"/"+url.QueryEscape(name))
	if err != nil {
		redBold(response)
	} else {
		greenBold(response)
	}

}
