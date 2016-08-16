package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type IPResponse struct {
	Title string
	Error string
	IPs   []IP
}

type IP struct {
	ID         int
	IP         string
	Network    string
	Mask       string
	Gateway    string
	Attached   bool
	InstanceID int
}

type Ips struct{}

// List returns the set of IPs in user's pool
func (ips *Ips) List() {
	url := APIBase + "/ippool/list"
	ir := IPResponse{}
	err := cli.GetJSON(url, &ir)

	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ir.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("IP"),
			greenBold("Network"), greenBold("Mask"), greenBold("Gateway"),
			greenBold("Attached"), greenBold("InstanceID")})

		// FIXME - auto-format
		for i := 0; i < len(ir.IPs); i++ {
			ipid := strconv.Itoa(ir.IPs[i].ID)
			att := strconv.FormatBool(ir.IPs[i].Attached)
			iid := strconv.Itoa(ir.IPs[i].InstanceID)

			table.Append([]string{ipid,
				ir.IPs[i].IP,
				ir.IPs[i].Network,
				ir.IPs[i].Mask,
				ir.IPs[i].Gateway,
				att,
				iid})
		}

		table.Render()

	}

}

// Attach attaches an ip to an instance
func (ips *Ips) Attach(ipv4 string, domain string) {
	//FIXME
	var ip struct {
		IP     string `json:"IP"`
		Domain string `json:"Domain"`
	}
	ip.IP = ipv4
	ip.Domain = domain

	b, err := json.Marshal(ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := APIBase + "/ip/connect"
	ir := IPResponse{}
	err = cli.PostJSON(b, url, &ir)

	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ir.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("IP"),
			greenBold("Network"), greenBold("Mask"), greenBold("Gateway"),
			greenBold("Attached"), greenBold("InstanceID")})

		// FIXME - auto-format
		for i := 0; i < len(ir.IPs); i++ {
			ipid := strconv.Itoa(ir.IPs[i].ID)
			att := strconv.FormatBool(ir.IPs[i].Attached)
			iid := strconv.Itoa(ir.IPs[i].InstanceID)

			table.Append([]string{ipid,
				ir.IPs[i].IP,
				ir.IPs[i].Network,
				ir.IPs[i].Mask,
				ir.IPs[i].Gateway,
				att,
				iid})
		}

		table.Render()

	}

}

// Detach detaches an ip to an instance
func (ips *Ips) Detach(ipv4 string) {
	//FIXME
	var ip struct {
		IP     string `json:"IP"`
		Domain string `json:"Domain"`
	}
	ip.IP = ipv4

	b, err := json.Marshal(ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := APIBase + "/ip/disconnect"
	ir := IPResponse{}
	err = cli.PostJSON(b, url, &ir)

	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ir.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("IP"),
			greenBold("Network"), greenBold("Mask"), greenBold("Gateway"),
			greenBold("Attached"), greenBold("InstanceID")})

		// FIXME - auto-format
		for i := 0; i < len(ir.IPs); i++ {
			ipid := strconv.Itoa(ir.IPs[i].ID)
			att := strconv.FormatBool(ir.IPs[i].Attached)
			iid := strconv.Itoa(ir.IPs[i].InstanceID)

			table.Append([]string{ipid,
				ir.IPs[i].IP,
				ir.IPs[i].Network,
				ir.IPs[i].Mask,
				ir.IPs[i].Gateway,
				att,
				iid})
		}

		table.Render()

	}

}

// Request requests for an ip to be added to a user's pool
func (ips *Ips) Request() {
	url := APIBase + "/ippool/request"
	ir := IPResponse{}
	err := cli.GetJSON(url, &ir)

	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ir.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("IP"),
			greenBold("Network"), greenBold("Mask"), greenBold("Gateway"),
			greenBold("Attached"), greenBold("InstanceID")})

		// FIXME - auto-format
		for i := 0; i < len(ir.IPs); i++ {
			ipid := strconv.Itoa(ir.IPs[i].ID)
			att := strconv.FormatBool(ir.IPs[i].Attached)
			iid := strconv.Itoa(ir.IPs[i].InstanceID)

			table.Append([]string{ipid,
				ir.IPs[i].IP,
				ir.IPs[i].Network,
				ir.IPs[i].Mask,
				ir.IPs[i].Gateway,
				att,
				iid})
		}

		table.Render()

	}

}

// Release releases an ip from a user's pool
func (ips *Ips) Release(ipv4 string) {

	//FIXME
	var ip struct {
		IP string `json:"IP"`
	}
	ip.IP = ipv4

	b, err := json.Marshal(ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := APIBase + "/ippool/return"
	ir := IPResponse{}
	err = cli.PostJSON(b, url, &ir)

	if err != nil {
		fmt.Println(redBold(err.Error()))
	} else {
		fmt.Println(greenBold(ir.Title))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		// FIXME - override headers
		table.SetHeader([]string{greenBold("ID"), greenBold("IP"),
			greenBold("Network"), greenBold("Mask"), greenBold("Gateway"),
			greenBold("Attached"), greenBold("InstanceID")})

		// FIXME - auto-format
		for i := 0; i < len(ir.IPs); i++ {
			ipid := strconv.Itoa(ir.IPs[i].ID)
			att := strconv.FormatBool(ir.IPs[i].Attached)
			iid := strconv.Itoa(ir.IPs[i].InstanceID)

			table.Append([]string{ipid,
				ir.IPs[i].IP,
				ir.IPs[i].Network,
				ir.IPs[i].Mask,
				ir.IPs[i].Gateway,
				att,
				iid})
		}

		table.Render()

	}
}
