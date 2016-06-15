package main

import (
	"encoding/json"
	"log"
	"os"
)

type Backups struct{}

func (backups *Backups) Save(domain string, name string) {
	instance := &Instance{}
	instance.Domain = domain
	instance.Name = name

	b, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(b, savebackupURL)
	if err != nil {
		log.Println(redBold(response))
	} else {
		log.Println(greenBold(response))
	}
}

func (backups *Backups) Restore(domain string, name string) {
	instance := &Instance{}
	instance.Domain = domain
	instance.Name = name
	b, err := json.Marshal(instance)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, restorebackupURL)
	if err != nil {
		log.Println(redBold(response))
	} else {
		log.Println(greenBold(response))
	}
}

func (backups *Backups) List() {
	response, err := cli.Postit(nil, listbackupsURL)
	if err != nil {
		log.Println(redBold(response))
	} else {
		log.Println(greenBold(response))
	}
}
