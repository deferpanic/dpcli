package main

import (
	"encoding/json"
	"log"
	"os"
)

var (
	// savebackupURL is the url to save a backup for running rumprun image
	savebackupURL = backupURL + "/save"

	// restorebackupURL is the url to restore a backup for running rumprun image
	restorebackupURL = backupURL + "/restore"

	// listbackupsURL is the url to list all rumprun image backups
	listbackupsURL = backupURL + "/list"
)

type Backups struct{}

func (backups *Backups) Save(name string, domain string) {
	instance := &OldInstance{}
	instance.Name = name
	instance.Domain = domain

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

func (backups *Backups) Restore(name string, domain string) {
	instance := &OldInstance{}

	instance.Name = name
	instance.Domain = domain

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
