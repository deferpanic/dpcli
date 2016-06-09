package main

import (
	"encoding/json"
	"flag"
	"github.com/deferpanic/dpcli/middleware"
	"log"
	"os"
)

var (
	scaleupPtr   = flag.Bool("scaleup", false, "Scale up execution of image")
	countPtr     = flag.Int("count", 1, "Number of image instances to launch")
	forcePtr     = flag.Bool("force", false, "Ignore errors during command execution")
	scaledownPtr = flag.Bool("scaledown", false, "Scale down execution of image")
	domainPtr    = flag.String("domain", "", "Domain name of image instance")
	runlogPtr    = flag.Bool("runlog", false, "View execution log of image instance")
	showPtr      = flag.Bool("show", false, "Show image instances")
	pausePtr     = flag.Bool("pause", false, "Stop image instance")
	resumePtr    = flag.Bool("resume", false, "Start image instance")
)

// processInstances process instance operations
func processInstances(cli middleware.RumpRunCLIInterface) (response string, executed bool, err error) {
	var b []byte

	executed = false
	if *scaleupPtr {
		image := &Image{}
		if *namePtr == "" {
			log.Println("Please provide image name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		image.Name = *namePtr
		if *countPtr < 1 {
			log.Println("Number of image instances to launch can't be less than 1")
			flag.PrintDefaults()
			os.Exit(1)
		}
		image.Count = *countPtr
		image.Force = *forcePtr
		b, err = json.Marshal(image)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, scaleupURL)
	}

	if *scaledownPtr {
		if *namePtr == "" && *domainPtr == "" {
			log.Println("Please provide image name or image instance domain name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance := &Instance{}
		instance.Name = *namePtr
		instance.Domain = *domainPtr
		instance.Force = *forcePtr
		b, err = json.Marshal(instance)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, scaledownURL)
	}

	if *runlogPtr {
		instance := &Instance{}
		if *domainPtr == "" {
			log.Println("Please provide image instance domain name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance.Domain = *domainPtr
		b, err = json.Marshal(instance)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, runlogURL)
	}

	if *showPtr {
		image := &Image{}
		if *namePtr == "" {
			log.Println("Please provide image name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		image.Name = *namePtr
		b, err = json.Marshal(image)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, showURL)
	}

	if *pausePtr {
		instance := &Instance{}
		if *domainPtr == "" {
			log.Println("Please provide image instance domain name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance.Domain = *domainPtr
		b, err = json.Marshal(instance)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, pauseURL)
	}

	if *resumePtr {
		instance := &Instance{}
		if *domainPtr == "" {
			log.Println("Please provide image instance domain name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance.Domain = *domainPtr
		b, err = json.Marshal(instance)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, resumeURL)
	}

	return response, executed, err
}
