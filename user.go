// Package main implements access to the deferpanic rumprun api for users.
package main

import (
	"flag"
	"fmt"
	"github.com/deferpanic/dpcli/middleware"
	"io/ioutil"
	"log"
	"os"
)

var (
	tokenPtr       = flag.String("token", "", "Token for making API calls. Must be provided")
	interactivePtr = flag.Bool("n", false, "non-interactive mode for scripting")
)

func main() {
	var err error

	flag.Parse()

	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.dprc")
	if err != nil {
		fmt.Println(err)
	}
	token := string(dat)

	if !*newPtr && !*displayPtr && !*makelogPtr && !*downloadPtr && !*uploadPtr &&
		!*scaleupPtr && !*scaledownPtr && !*runlogPtr && !*showPtr && !*pausePtr && !*resumePtr &&
		!*savebackupPtr && !*restorebackupPtr && !*listbackupsPtr &&
		!*newresPtr && !*listresPtr &&
		!*statusPtr && !*versionPtr && !*languagesPtr && !*addonsPtr && !*builtinsPtr {
		log.Println("Please chose one of available commands")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if token == "" {
		token = *tokenPtr
	}

	if token == "" && !(*statusPtr || *versionPtr || *languagesPtr || *addonsPtr || *builtinsPtr) {
		log.Println("Please provide API token")
		flag.PrintDefaults()
		os.Exit(1)
	}

	cli := middleware.NewRumpRunCLIImplementation(token)

	var response string

	executed := false
	if !executed {
		response, executed, err = processImages(cli)
	}

	if !executed {
		response, executed, err = processInstances(cli)
	}

	if !executed {
		response, executed, err = processBackups(cli)
	}

	if !executed {
		response, executed, err = processResources(cli)
	}

	if !executed {
		response, executed, err = processSystem(cli)
	}

	if !*interactivePtr {

		if err != nil {
			fmt.Println(redBold(response))
		} else {
			fmt.Println(greenBold(response))
		}

	} else {
		fmt.Println(response)
	}
}
