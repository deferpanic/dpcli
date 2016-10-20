package api

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

var (
	// scaleupURL is the url to add another instance of a project
	scaleupURL = instanceURL + "/scaleup"

	// scaledownURL is the url to scale down a project
	// this will currently scale down every instance
	scaledownURL = instanceURL + "/scaledown"

	// runlogURL is the url to view the log of a running project
	runlogURL = instanceURL + "/log"

	// showURL is the url to list the instances
	showURL = instanceURL + "/show"

	// pauseURL is the url to stop a running project
	pauseURL = instanceURL + "/pause"

	// resumeURL is the url to start a project
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
	response, err := Cli.Postit(b, scaleupURL)
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
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

	response, err := Cli.Postit(b, runlogURL)
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
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
	err := Cli.GetJSON(url, &ir)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
	} else {
		fmt.Println(GreenBold(ir.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{GreenBold("ID"), GreenBold("Domain"),
			GreenBold("NetworkID"), GreenBold("Memory"), GreenBold("Disk"),
			GreenBold("Status"), GreenBold("Running"), GreenBold("StartedAt")})

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

	response, err := Cli.Postit(b, pauseURL)
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
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

	response, err := Cli.Postit(b, resumeURL)
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
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

	response, err := Cli.Postit(b, scaleupURL)
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
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

	response, err := Cli.Postit(b, scaledownURL)
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}

}
