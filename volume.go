package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (

	// URL is the url for volume management
	volumeURL = APIBase + "/volume"

	// newvolumeURL is the url to add a volume for rumprun image
	newvolumeURL = volumeURL + "/new"

	// removevolumeURL is the url to remove a volume from rumprun image
	removevolumeURL = volumeURL + "/remove"

	// connectvolumeURL is the url to connect a volume to rumprun image
	connectvolumeURL = volumeURL + "/connect"

	// disconnectvolumeURL is the url to disconnect a volume from
	// rumprun image
	disconnectvolumeURL = volumeURL + "/disconnect"

	// listvolumesURL is the url to list all rumprun image volumes
	listvolumesURL = volumeURL + "/list"

	// getvolumeURL is the url to download rumprun image volume
	getvolumeURL = volumeURL + "/get"

	// putvolumeURL is the url to upload rumprun image volume
	putvolumeURL = volumeURL + "/put"
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
	volume.Owner = name

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

func (volumes *Volumes) Attach(name string, domain string) {
	volume := &Volume{}
	volume.Name = name
	volume.Domain = domain

	b, err := json.Marshal(volume)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(b, connectvolumeURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}

func (volumes *Volumes) Detach(name string, domain string) {
	volume := &Volume{}
	volume.Name = name
	volume.Domain = domain

	b, err := json.Marshal(volume)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, disconnectvolumeURL)
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}
