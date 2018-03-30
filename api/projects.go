package api

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Projects struct{}

// New creates a new project
func (projects *Projects) New(name string, language string, compiler string, source string, script string) {

	image := &Image{}
	image.Name = name
	image.Language = language
	image.Compiler = compiler
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

	response, err := Cli.Postit(b, APIBase+"/image/new")
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
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

	response, err := Cli.Postit(b, APIBase+"/image/remove")
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}
}

// NewFromImage creates a new project from an image
func (projects *Projects) NewFromImage(name string, imagePath string) {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response, err := Cli.Postit(data, putURL+"/"+url.QueryEscape(name))
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
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
	err := Cli.GetJSON(APIBase+"/image/display", &pr)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
	} else {
		fmt.Println(GreenBold(pr.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)

		table.SetHeader([]string{GreenBold("ID"), GreenBold("Name"),
			GreenBold("Buildable"), GreenBold("Language"), GreenBold("Source"),
			GreenBold("BuildStatus"), GreenBold("Filename"), GreenBold("Addon")})

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

type ProjectLogResponse struct {
	Error string
	Body  string
}

// Log shows the latest build log for your project by name
func (projects *Projects) Log(name string) {
	plr := ProjectLogResponse{}
	err := Cli.GetJSON(APIBase+"/builds/"+name+"/latest.json", &plr)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
	} else {
		fmt.Println(GreenBold(plr.Body))
	}
}

// Download a project root image
func (projects *Projects) Download(name string, path string) error {
	image := &Image{}
	image.Name = name

	b, err := json.Marshal(image)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = Cli.GrabFile(b, imageURL+"/get", path); err != nil {
		return err
	}

	return nil
}

// DownloadCommunity downloads a community kernel by project name and
// user name
// the project that holds this kernel must be public
func (projects *Projects) DownloadCommunity(name string, user string, path string) error {
	if err := Cli.GrabFile(nil, APIBase+"/kernels/download/"+user+"/"+name, path); err != nil {
		return err
	}

	return nil
}

// Upload uploads a project
func (projects *Projects) Upload(name string, binary string) {
	data, err := ioutil.ReadFile(binary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response, err := Cli.Postit(data, putURL+"/"+url.QueryEscape(name))
	if err != nil {
		fmt.Println(RedBold(response))
	} else {
		fmt.Println(GreenBold(response))
	}

}

type ManifestVolume struct {
	Id    int
	File  string
	Mount string
}

type ManifestProcess struct {
	Memory    int
	Kernel    string
	Multiboot bool
	Hash      string
	Cmdline   string
	Env       string
	Volumes   []ManifestVolume
}

type Manifest struct {
	Processes []ManifestProcess
}

// Manifest is a json representation for your project
// useful for running things locally
func (projects *Projects) Manifest(name string) error {
	m := Manifest{}
	err := Cli.GetJSON(APIBase+"/projects/manifest/"+name, &m)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
		return err
	}

	// dupey-dupe
	js, err := json.Marshal(m)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
		return err
	}

	if strings.Contains(name, "/") {
		name = strings.Replace(name, "/", "_", -1)
	}

	err = ioutil.WriteFile(name+".manifest", js, 0644)
	if err != nil {
		fmt.Println(RedBold(err.Error()))
		return err
	}

	return nil
}

// Temporary stuff for `virgo` refactoring only
//
func LoadManifest(name string) (Manifest, error) {
	m := Manifest{}

	err := Cli.GetJSON(APIBase+"/projects/manifest/"+name, &m)
	if err != nil {
		return Manifest{}, err
	}

	return m, err
}
