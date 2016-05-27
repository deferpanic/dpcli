package main

import (
	"encoding/json"
	"flag"
	"github.com/deferpanic/dpcli/middleware"
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
	binaryPtr   = flag.String("binary", "", "Path to image binary")
	displayPtr  = flag.Bool("display", false, "Display all images")
	makelogPtr  = flag.Bool("makelog", false, "View making log of image")
	downloadPtr = flag.Bool("download", false, "Download image binary")
	uploadPtr   = flag.Bool("upload", false, "Upload image binary")
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
			if *binaryPtr != "" {
				image.Data, err = ioutil.ReadFile(*binaryPtr)
				if err != nil {
					log.Println(err)
					os.Exit(1)
				}
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
	if *displayPtr {
		executed = true
		response, err = cli.Postit(b, displayURL)
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
	if *downloadPtr {
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
		response, err = cli.Postit(b, getURL)
	}
	if *uploadPtr {
		image := &Image{}
		if *namePtr == "" {
			log.Println("Please provide image name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		image.Name = *namePtr
		if *binaryPtr == "" {
			log.Println("Please provide path to image binary")
			flag.PrintDefaults()
			os.Exit(1)
		}
		image.Data, err = ioutil.ReadFile(*binaryPtr)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		b, err = json.Marshal(image)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, putURL)
	}

	return response, executed, err
}
