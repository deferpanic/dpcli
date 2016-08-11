package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

var (
	// scaleupURL is the url to scale up execution of rumprun image
	scaleupURL = instanceURL + "/scaleup"

	// scaledownURL is the url to scale down execution of rumprun image
	scaledownURL = instanceURL + "/scaledown"

	// runlogURL is the url to view execution log of rumprun image
	runlogURL = instanceURL + "/log"

	// showURL is the url to show running rumprun images
	showURL = instanceURL + "/show"

	// pauseURL is the url to stop existing rumprun image
	pauseURL = instanceURL + "/pause"

	// resumeURL is the url to start existing rumprun image
	resumeURL = instanceURL + "/resume"
)

// DEPRECATED
// OldInstance is the base struct for management of instances
type OldInstance struct {
	Name   string `json:"Name"`
	Domain string `json:"Domain"`
	Force  bool   `json:"Force"`
}

type Instance struct {
	ID        int
	Domain    string
	NetworkID int
	Memory    int
	Disk      int64
	Status    string
	Running   bool
	StartedAt time.Time
}

type InstancesResponse struct {
	Title     string
	Error     string
	Instances []Instance
}

type Instances struct{}

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
	instance := &OldInstance{}
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

// List returns the set of running instances for project name
// if name is empty it returns all running instances for user
func (instances *Instances) List(name string) {
	url := ""
	if name != "" {
		url = APIBase + "/instances/list/" + name
	} else {
		url = APIBase + "/instances/list"
	}

	ir := InstancesResponse{}
	err := cli.GetJSON(url, &ir)
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ir.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("Domain"),
			greenBold("NetworkID"), greenBold("Memory"), greenBold("Disk"),
			greenBold("Status"), greenBold("Running"), greenBold("StartedAt")})

		// FIXME - auto-format
		for i := 0; i < len(ir.Instances); i++ {
			sid := strconv.Itoa(ir.Instances[i].ID)
			nid := strconv.Itoa(ir.Instances[i].NetworkID)
			mem := strconv.Itoa(ir.Instances[i].Memory)
			run := strconv.FormatBool(ir.Instances[i].Running)

			table.Append([]string{sid,
				ir.Instances[i].Domain,
				nid,
				mem,
				formatSz(ir.Instances[i].Disk),
				ir.Instances[i].Status,
				run,
				ir.Instances[i].StartedAt.String()})
		}

		table.Render()

	}

}

func (instances *Instances) Pause(domain string) {
	instance := &OldInstance{}
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
	instance := &OldInstance{}
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
	instance := &OldInstance{}
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
