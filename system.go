package main

import (
	"flag"
	"github.com/deferpanic/dpcli/middleware"
)

var (
	languagesPtr = flag.Bool("languages", false, "Print supported programming languages")
	addonsPtr    = flag.Bool("addons", false, "Print addons available")
	builtinsPtr  = flag.Bool("builtins", false, "Print builtins available")
	statusPtr    = flag.Bool("status", false, "Print current system status")
	versionPtr   = flag.Bool("version", false, "Display current API version")
)

// processSystem process system operations
func processSystem(cli middleware.RumpRunCLIInterface) (response string, executed bool, err error) {
	var b []byte

	executed = false
	if *languagesPtr {
		executed = true
		response, err = cli.Postit(b, languagesURL)
	}
	if *addonsPtr {
		executed = true
		response, err = cli.Postit(b, addonsURL)
	}
	if *builtinsPtr {
		executed = true
		response, err = cli.Postit(b, builtinsURL)
	}
	if *statusPtr {
		executed = true
		response, err = cli.Postit(b, statusURL)
	}
	if *versionPtr {
		executed = true
		response = "\n" + "API version: " + APIVersion + "\n"
	}

	return response, executed, err
}
