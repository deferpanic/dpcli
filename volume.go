package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Volume is the base struct for management of volumes
type Volume struct {
	Name       string `json:"Name"`
	Owner      string `json:"Owner"`
	Domain     string `json:"Domain"`
	MountPoint string `json:"MountPoint"`
	Filename   string `json:"Filename"`
}

type Volumes struct{}

func (volumes *Volumes) ListByName(name string) {
	volume := &Volume{}
	volume.Name = name

	volumes.List(volume)
}

func (volumes *Volumes) ListByDomain(domain string) {
	volume := &Volume{}
	volume.Domain = domain

	volumes.List(volume)
}

func (volumes *Volumes) List(volume *Volume) {

	b, err := json.Marshal(volume)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, APIBase+"/volume/list")
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}
