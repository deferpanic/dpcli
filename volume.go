package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type Volume struct {
	ID         int
	Name       string
	MountPoint string
	Filename   string
	ProjectId  int
	Size       string
	InstanceId int
	Orphaned   bool
	Attached   bool
	Mutable    bool
	FileSystem string
	System     bool
	CreatedAt  string
}

type VolumesResponse struct {
	Title   string
	Error   string
	Volumes []Volume
}

//FIXME -- get rid of all this crap
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

	// putvolumeURL is the url to upload rumprun image volume
	putvolumeURL = volumeURL + "/put"
)

// OldVolume is the base struct for management of volumes
// DEPRECATED
type OldVolume struct {
	Id         int    `json:"ID"`
	Name       string `json:"Name"`
	Owner      string `json:"Owner"`
	Domain     string `json:"Domain"`
	MountPoint string `json:"MountPoint"`
	Filename   string `json:"Filename"`
}

type Volumes struct{}

// ListByName lists volumes that are within a project by name
func (volumes *Volumes) ListByName(name string) {
	volume := &OldVolume{}
	volume.Owner = name

	volumes.List(volume)
}

// ListByDomain lists volumes attached to a given domain
func (volumes *Volumes) ListByDomain(domain string) {
	volume := &OldVolume{}
	volume.Domain = domain

	volumes.List(volume)
}

// List lists volumes
func (volumes *Volumes) List(volume *OldVolume) {
	b, err := json.Marshal(volume)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	vr := VolumesResponse{}
	err = cli.PostJSON(b, APIBase+"/volume/list", &vr)
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(vr.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)

		table.SetHeader([]string{greenBold("ID"), greenBold("Name"),
			greenBold("Mount Point"), greenBold("Filename"), greenBold("ProjectID"),
			greenBold("Size"), greenBold("InstanceId"), greenBold("Orphaned"),
			greenBold("Attached"), greenBold("Mutable"), greenBold("FileSystem"),
			greenBold("System"), greenBold("CreatedAt"),
		})

		for i := 0; i < len(vr.Volumes); i++ {
			sid := strconv.Itoa(vr.Volumes[i].ID)
			pid := strconv.Itoa(vr.Volumes[i].ProjectId)
			iid := strconv.Itoa(vr.Volumes[i].InstanceId)
			table.Append([]string{sid,
				vr.Volumes[i].Name,
				vr.Volumes[i].MountPoint,
				vr.Volumes[i].Filename,
				pid,
				vr.Volumes[i].Size,
				iid,
				strconv.FormatBool(vr.Volumes[i].Orphaned),
				strconv.FormatBool(vr.Volumes[i].Attached),
				strconv.FormatBool(vr.Volumes[i].Mutable),
				vr.Volumes[i].FileSystem,
				strconv.FormatBool(vr.Volumes[i].System),
				vr.Volumes[i].CreatedAt,
			})
		}

		table.Render()
	}

}

// Attach attaches a volume from a project
func (volumes *Volumes) Attach(name string, domain string) {
	volume := &OldVolume{}
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

// Detach detaches a volume from a project
func (volumes *Volumes) Detach(name string, domain string) {
	volume := &OldVolume{}
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

// Downloads a volume by id
// FIXME
func (volumes *Volumes) Download(id int) {

	err := cli.GrabFile(nil, volumeURL+"/get/"+strconv.Itoa(id), "vol"+strconv.Itoa(id))
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold("file saved"))
	}

}
