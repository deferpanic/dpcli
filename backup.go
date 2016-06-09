package main

import (
	"encoding/json"
	"flag"
	"github.com/deferpanic/dpcli/middleware"
	"log"
	"os"
)

var (
	savebackupPtr    = flag.Bool("savebackup", false, "Save backup of image instance. Must be paused before")
	restorebackupPtr = flag.Bool("restorebackup", false, "Restore backup for image instance. Must be paused before")
	listbackupsPtr   = flag.Bool("listbackups", false, "List all image backups")
)

// processBackups process backup operations
func processBackups(cli middleware.RumpRunCLIInterface) (response string, executed bool, err error) {
	var b []byte

	executed = false
	if *savebackupPtr {
		instance := &Instance{}
		if *domainPtr == "" {
			log.Println("Please provide image instance domain name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance.Domain = *domainPtr
		if *namePtr == "" {
			log.Println("Please provide image backup name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance.Name = *namePtr
		b, err = json.Marshal(instance)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, savebackupURL)

	}

	if *restorebackupPtr {
		instance := &Instance{}
		if *domainPtr == "" {
			log.Println("Please provide image instance domain name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance.Domain = *domainPtr
		if *namePtr == "" {
			log.Println("Please provide image backup name")
			flag.PrintDefaults()
			os.Exit(1)
		}
		instance.Name = *namePtr
		b, err = json.Marshal(instance)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		executed = true
		response, err = cli.Postit(b, restorebackupURL)
	}

	if *listbackupsPtr {
		executed = true
		response, err = cli.Postit(b, listbackupsURL)
	}

	return response, executed, err
}
