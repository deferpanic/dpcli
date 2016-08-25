package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Projects struct{}

// New creates a new project
func (projects *Projects) New(name string, language string, source string, script string) {

	image := &Image{}
	image.Name = name
	image.Language = language
	image.MakeBin = true
	image.Source = source
	image.SystemVolumes = true
	image.Buildable = true

	data, err := ioutil.ReadFile(script)
	if err == nil {
		image.Script = string(data)
	}

	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, APIBase+"/image/new")
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}

// Delete deletes a project
func (projects *Projects) Delete(name string) {

	image := &Image{}
	image.Name = name
	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(b, APIBase+"/image/remove")
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}
}

// NewFromImage creates a new project from an image
func (projects *Projects) NewFromImage(name string, imagePath string) {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(data, putURL+"/"+url.QueryEscape(name))
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}

type Project struct {
	ID          int
	Name        string
	Buildable   string
	Language    string
	Source      string
	BuildStatus string
	Filename    string
	Addon       string
	CreatedAt   time.Time
}

type ProjectsResponse struct {
	Title    string
	Error    string
	Projects []Project
}

// List lists all your projects
func (projects *Projects) List() {
	pr := ProjectsResponse{}
	err := cli.GetJSON(APIBase+"/image/display", &pr)
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(pr.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)

		table.SetHeader([]string{greenBold("ID"), greenBold("Name"),
			greenBold("Buildable"), greenBold("Language"), greenBold("Source"),
			greenBold("BuildStatus"), greenBold("Filename"), greenBold("Addon")})

		for i := 0; i < len(pr.Projects); i++ {
			sid := strconv.Itoa(pr.Projects[i].ID)

			table.Append([]string{sid,
				pr.Projects[i].Name,
				pr.Projects[i].Buildable,
				pr.Projects[i].Language,
				pr.Projects[i].Source,
				pr.Projects[i].BuildStatus,
				pr.Projects[i].Filename,
				pr.Projects[i].Addon,
			})
		}

		table.Render()
	}

}

// Log shows the log output for your project
func (projects *Projects) Log(name string) {
	image := &Image{}

	image.Name = name
	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response, err := cli.Postit(b, APIBase+"/image/makelog")
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}

// Download a project root image
func (projects *Projects) Download(name string) {

	image := &Image{}
	image.Name = name
	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cli.GrabFile(b, imageURL+"/get", name)
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold("file saved"))
	}

}

// Upload uploads a project
func (projects *Projects) Upload(name string, binary string) {
	data, err := ioutil.ReadFile(binary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := cli.Postit(data, putURL+"/"+url.QueryEscape(name))
	if err != nil {
		fmt.Println(redBold(response))
	} else {
		fmt.Println(greenBold(response))
	}

}
