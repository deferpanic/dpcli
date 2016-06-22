// Package main implements access to the deferpanic unikernel IaaS API
// for users.
package main

import (
	"fmt"
	"github.com/deferpanic/dpcli/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app = kingpin.New("dpcli", "Tooling to interact with DeferPanic IaaS")

	token       = app.Flag("token", "Token").String()
	interactive = app.Flag("interactive", "Disable interactive mode .").Bool()

	projectsCommand     = app.Command("projects", "Projects.")
	projectsNewCommand  = projectsCommand.Command("new", "Create a new project.")
	projectsNewName     = projectsNewCommand.Arg("name", "Project name.").Required().String()
	projectsNewLanguage = projectsNewCommand.Arg("language", "Project language.").Required().String()
	projectsNewSource   = projectsNewCommand.Arg("source", "Project source.").Required().String()
	projectsNewScript   = projectsNewCommand.Arg("script", "Project script.").String()

	projectsDeleteCommand = projectsCommand.Command("delete", "Delete a project.")
	projectsDeleteName    = projectsDeleteCommand.Arg("name", "Project name.").Required().String()

	projectsDownloadCommand = projectsCommand.Command("download", "Download image.")
	projectsUploadCommand   = projectsCommand.Command("upload", "Upload image.")
	projectsUploadBinary    = projectsUploadCommand.Arg("binary", "Image binary path.").Required().String()

	projectsListCommand = projectsCommand.Command("list", "List Projects.")
	projectsLogCommand  = projectsCommand.Command("log", "View Project Build Log")
	projectsLogName     = projectsLogCommand.Arg("name", "Project name.").Required().String()

	instancesCommand    = app.Command("instances", "Instances.")
	instancesNewCommand = instancesCommand.Command("new", "Create a new instance.")
	instancesNewName    = instancesNewCommand.Arg("name", "Project name.").Required().String()

	instancesLogCommand = instancesCommand.Command("log", "Show log of instance.")
	instancesLogName    = instancesLogCommand.Arg("name", "Instance name.").Required().String()

	instancesListCommand = instancesCommand.Command("list", "List instances attached to project.")
	instancesListName    = instancesListCommand.Arg("name", "Project name.").String()

	instancesPauseCommand = instancesCommand.Command("pause", "Pause instance.")
	instancesPauseName    = instancesPauseCommand.Arg("name", "Project name.").Required().String()

	instancesResumeCommand = instancesCommand.Command("resume", "Resume Instance.")
	instancesResumeName    = instancesResumeCommand.Arg("name", "Project name.").Required().String()

	instancesScaleUpCommand = instancesCommand.Command("scaleup", "ScaleUp Instance.")
	instancesScaleUpName    = instancesScaleUpCommand.Arg("name", "Project name.").Required().String()

	instancesScaleDownCommand = instancesCommand.Command("scaledown", "ScaleDown Instance.")
	instancesScaleDownName    = instancesScaleDownCommand.Arg("name", "Project name.").Required().String()
	instancesScaleDownDomain  = instancesScaleDownCommand.Arg("domain", "Domain").Required().String()

	volumesCommand     = app.Command("volumes", "Volumes.")
	volumesListCommand = volumesCommand.Command("list", "List volumes")
	volumesListName    = volumesListCommand.Flag("name", "Project name.").String()
	volumesListDomain  = volumesListCommand.Flag("domain", "Domain.").String()

	volumesCreateCommand = volumesCommand.Command("create", "Create volume")
	volumesShowCommand   = volumesCommand.Command("show", "Show Volume")
	volumesUpdateCommand = volumesCommand.Command("update", "Update Volume")
	volumesDeleteCommand = volumesCommand.Command("delete", "Delete Volume")

	volumesAttachCommand = volumesCommand.Command("attach", "Attach Volume")
	volumesAttachName    = volumesAttachCommand.Flag("name", "Project name.").String()
	volumesAttachDomain  = volumesAttachCommand.Flag("domain", "Domain.").String()

	volumesDetachCommand = volumesCommand.Command("detach", "Detach Volume")
	volumesDetachName    = volumesDetachCommand.Flag("name", "Project name.").String()
	volumesDetachDomain  = volumesDetachCommand.Flag("domain", "Domain.").String()

	volumesDownloadCommand = volumesCommand.Command("download", "Download Volume")
	volumesUploadCommand   = volumesCommand.Command("upload", "Upload Volume")

	backupsCommand     = app.Command("backups", "Backups.")
	backupsSaveCommand = backupsCommand.Command("save", "Save backup of image instance.")
	backupsSaveName    = backupsSaveCommand.Arg("name", "Instance name.").Required().String()
	backupsSaveDomain  = backupsSaveCommand.Arg("domain", "Domain name.").Required().String()

	backupsRestoreCommand = backupsCommand.Command("restore", "Restore an image instance.")
	backupsRestoreName    = backupsRestoreCommand.Arg("name", "Instance name.").Required().String()
	backupsRestoreDomain  = backupsRestoreCommand.Arg("domain", "Domain name.").Required().String()

	backupsListCommand = backupsCommand.Command("list", "List available backups")

	languagesCommand = app.Command("languages", "Languages.")

	resourcesCommand    = app.Command("resources", "Resources.")
	resourcesNewCommand = resourcesCommand.Command("new", "Add a resource.")
	resourcesNewName    = resourcesNewCommand.Arg("name", "name?").Required().String()
	resourcesNewOwner   = resourcesNewCommand.Arg("builtin", "builtin").Required().String()
	resourcesNewBuiltin = resourcesNewCommand.Arg("owner", "owner?").Required().String()

	resourcesListCommand = resourcesCommand.Command("list", "List available resources.")
	resourcesListName    = resourcesListCommand.Arg("name", "name.").Required().String()

	addonsCommand = app.Command("addons", "Addons.")

	builtinsCommand = app.Command("builtins", "Builtins.")

	status = app.Command("status", "Show Status.")
)

// fixme
// should only have to be called once..
func setToken() {
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.dprc")
	if err != nil {
		fmt.Println(redBold("you can stick your token in ~/.dprc"))
	}
	dtoken := string(dat)

	if dtoken == "" {
		dtoken = *token
	}

	if dtoken == "" {
		redBold("no token")
		os.Exit(1)
	}

	cli = middleware.NewRumpRunCLIImplementation(dtoken)
}

var cli *middleware.RumpRunCLIImplementation

func main() {

	kingpin.Version("0.0.1")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "projects new":
		setToken()
		projects := &Projects{}
		projects.New(*projectsNewName, *projectsNewLanguage, *projectsNewSource, *projectsNewScript)
	case "projects delete":
		setToken()
		projects := &Projects{}
		projects.Delete(*projectsDeleteName)
	case "projects list":
		setToken()
		projects := &Projects{}
		projects.List()
	case "projects log":
		setToken()
		projects := &Projects{}
		projects.Log(*projectsLogName)
	case "instances scaleup":
		setToken()
		instances := &Instances{}
		instances.ScaleUp(*instancesScaleUpName)
	case "instances scaledown":
		setToken()
		instances := &Instances{}
		instances.ScaleDown(*instancesScaleDownName, *instancesScaleDownDomain)
	case "instances new":
		setToken()
		instances := &Instances{}
		instances.New(*instancesNewName)
	case "instances log":
		setToken()
		instances := &Instances{}
		instances.Log(*instancesLogName)
	case "instances list":
		setToken()
		instances := &Instances{}
		instances.List(*instancesListName)
	case "instances pause":
		setToken()
		instances := &Instances{}
		instances.Pause(*instancesPauseName)
	case "instances resume":
		setToken()
		instances := &Instances{}
		instances.Resume(*instancesResumeName)
	case "volumes list":
		setToken()
		volumes := &Volumes{}
		if *volumesListName != "" {
			volumes.ListByName(*volumesListName)
		}
		if *volumesListDomain != "" {
			volumes.ListByDomain(*volumesListDomain)
		}
	case "volumes create":
		setToken()
	case "volumes show":
		setToken()
	case "volumes update":
		setToken()
	case "volumes delete":
		setToken()
	case "volumes attach":
		setToken()
		volumes := &Volumes{}
		volumes.Attach(*volumesAttachName, *volumesAttachDomain)
	case "volumes detach":
		setToken()
		volumes := &Volumes{}
		volumes.Detach(*volumesAttachName, *volumesAttachDomain)
	case "volumes download":
		setToken()
	case "volumes upload":
		setToken()
	case "backups list":
		setToken()
		backups := &Backups{}
		backups.List()
	case "backups save":
		setToken()
		backups := &Backups{}
		backups.Save(*backupsSaveName, *backupsSaveDomain)
	case "backups restore":
		setToken()
		backups := &Backups{}
		backups.Restore(*backupsRestoreName, *backupsRestoreDomain)
	case "languages":
		setToken()
		languages := &Languages{}
		languages.List()
	case "resources create":
		setToken()
		resources := &Resources{}
		resources.New(*resourcesNewName, *resourcesNewOwner, *resourcesNewBuiltin)
	case "resources list":
		setToken()
		resources := &Resources{}
		resources.List(*resourcesListName)
	case "addons":
		setToken()
		addons := &Addons{}
		addons.List()
	case "status":
		setToken()
		status := &Status{}
		status.Show()
	}
}
