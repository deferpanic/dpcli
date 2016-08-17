package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type Addons struct{}

type Addon struct {
	ID          int
	Name        string
	Addon       string
	BuildStatus string
	CreatedAt   string
}

type AddonsResponse struct {
	Title  string
	Error  string
	Addons []Addon
}

type AddonAvailable struct {
	ID          int
	Name        string
	Description string
}

type AddonsAvailableResponse struct {
	Title           string
	Error           string
	AddonsAvailable []AddonAvailable
}

// Available lists the addons available for the user to provision
func (addons *Addons) Available() {

	ar := AddonsAvailableResponse{}
	err := cli.GetJSON(APIBase+"/addons/available", &ar)
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ar.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("Name"),
			greenBold("Description")})

		// FIXME - auto-format
		for i := 0; i < len(ar.AddonsAvailable); i++ {
			sid := strconv.Itoa(ar.AddonsAvailable[i].ID)

			table.Append([]string{sid,
				ar.AddonsAvailable[i].Name,
				ar.AddonsAvailable[i].Description})
		}

		table.Render()

	}

}

// List lists the addons provisioned to the user
func (addons *Addons) List() {
	ar := AddonsResponse{}
	err := cli.GetJSON(APIBase+"/addons/list", &ar)
	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ar.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("Name"),
			greenBold("Addon"), greenBold("BuildStatus"), greenBold("Created At")})

		// FIXME - auto-format
		for i := 0; i < len(ar.Addons); i++ {
			sid := strconv.Itoa(ar.Addons[i].ID)

			table.Append([]string{sid,
				ar.Addons[i].Name,
				ar.Addons[i].Addon,
				ar.Addons[i].BuildStatus,
				ar.Addons[i].CreatedAt})
		}

		table.Render()

	}

}

func (addons *Addons) New() {
}

func (addons *Addons) Delete() {
}
