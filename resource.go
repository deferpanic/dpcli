package main

import (
	"encoding/json"
	"flag"
	"github.com/deferpanic/dpcli/middleware"
	"log"
	"os"
)

var (
	newresPtr  = flag.Bool("newres", false, "Create new resource for image")
	resnamePtr = flag.String("resname", "", "Name of resource")
	builtinPtr = flag.String("builtin", "", "Builtin of resource")
	listresPtr = flag.Bool("listres", false, "List all image resources")
)

// processResources process resource operations
func processResources(cli middleware.RumpRunCLIInterface) (response string, executed bool, err error) {
	var b []byte

	executed = false
	if *newresPtr {
		resource := &Resource{}
		if *resnamePtr == "" {
			log.Println("Please provide resource image name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		resource.Name = *resnamePtr
		if *namePtr == "" {
			log.Println("Please provide image name for resource")
			flag.PrintDefaults()
			os.Exit(1)
		}
		resource.Owner = *namePtr
		if *builtinPtr == "" {
			log.Println("Please provide builtin name for resource")
			flag.PrintDefaults()
			os.Exit(1)
		}
		resource.Builtin = *builtinPtr
		b, err = json.Marshal(resource)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, newresourceURL)
	}
	if *listresPtr {
		resource := &Resource{}
		if *namePtr == "" {
			log.Println("Please provide image name for resource")
			flag.PrintDefaults()
			os.Exit(1)
		}
		resource.Owner = *namePtr
		b, err = json.Marshal(resource)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, listresourcesURL)
	}

	return response, executed, err
}
