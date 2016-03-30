package main

import (
	"encoding/json"
	"flag"
	"github.com/deferpanic/rumpruncli/middleware"
	"io/ioutil"
	"log"
	"os"
)

var (
	newPtr      = flag.Bool("new", false, "Create new image")
	namePtr     = flag.String("name", "", "Name of image")
	sourcePtr   = flag.String("source", "", "Source of image")
	addonPtr    = flag.String("addon", "", "Addon for image")
	scriptPtr   = flag.String("script", "script.yml", "Script for image")
	languagePtr = flag.String("language", "", "Source language of image")
	makelogPtr  = flag.Bool("makelog", false, "View making log of image")
)

// processImages process image operations
func processImages(cli middleware.RumpRunCLIInterface) (response string, executed bool, err error) {
	var b []byte

	executed = false
	if *newPtr {
		image := &Image{}
		if *namePtr == "" {
			log.Println("Please provide image name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		image.Name = *namePtr
		if *sourcePtr == "" && *addonPtr == "" {
			log.Println("Please provide image source or image addon")
			flag.PrintDefaults()
			os.Exit(1)
		}
		if *sourcePtr != "" {
			image.Source = *sourcePtr
		} else if *addonPtr != "" {
			image.Addon = *addonPtr
		}

		data, err := ioutil.ReadFile(*scriptPtr)
		if err == nil {
			image.Script = string(data)
		}
		if *sourcePtr != "" {
			if *languagePtr != "" {
				image.Language = *languagePtr
			} else {
				log.Println("Please provide image source language")
				flag.PrintDefaults()
				os.Exit(1)
			}
		}
		b, err = json.Marshal(image)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, newURL)
	}
	if *makelogPtr {
		image := &Image{}
		if *namePtr == "" {
			log.Println("Please provide image name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		image.Name = *namePtr
		b, err = json.Marshal(image)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, makelogURL)
	}

	return response, executed, err
}
