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

	token          = app.Flag("token", "Token").String()
	interactive    = app.Flag("interactive", "Disable interactive mode .").Bool()
	versionCommand = app.Command("version", "Version")

	projectsCommand     = app.Command("projects", "Projects.")
	projectsNewCommand  = projectsCommand.Command("new", "Create a new project.")
	projectsNewName     = projectsNewCommand.Arg("name", "Project name.").Required().String()
	projectsNewLanguage = projectsNewCommand.Arg("language", "Project language.").Required().String()
	projectsNewSource   = projectsNewCommand.Arg("source", "Project source.").Required().String()
	projectsNewScript   = projectsNewCommand.Arg("script", "Project script.").String()

	projectsDeleteCommand = projectsCommand.Command("delete", "Delete a project.")
	projectsDeleteName    = projectsDeleteCommand.Arg("name", "Project name.").Required().String()

	projectsDownloadCommand = projectsCommand.Command("download", "Download image.")
	projectsDownloadName    = projectsDownloadCommand.Arg("name", "Project name.").Required().String()
	projectsUploadCommand   = projectsCommand.Command("upload", "Upload image.")
	projectsUploadBinary    = projectsUploadCommand.Arg("binary", "Image binary path.").Required().String()

	projectsListCommand = projectsCommand.Command("list", "List Projects.")
	projectsLogCommand  = projectsCommand.Command("log", "View Project Build Log")
	projectsLogName     = projectsLogCommand.Arg("name", "Project name.").Required().String()

	projectsManifestCommand = projectsCommand.Command("manifest", "Project manifest.")
	projectsManifestName    = projectsManifestCommand.Arg("name", "Project name.").Required().String()

	instancesCommand    = app.Command("instances", "Instances.")
	instancesNewCommand = instancesCommand.Command("new", "Create a new instance.")
	instancesNewName    = instancesNewCommand.Arg("name", "Project name.").Required().String()

	instancesLogCommand = instancesCommand.Command("log", "Show log of instance.")
	instancesLogName    = instancesLogCommand.Arg("name", "Instance name.").Required().String()

	instancesListCommand = instancesCommand.Command("list", "List instances attached to project.")
	instancesListName    = instancesListCommand.Arg("name", "Project name.").String()

	instancesPauseCommand = instancesCommand.Command("pause", "Pause instance.")
	instancesPauseName    = instancesPauseCommand.Arg("domain", "Instance domain").Required().String()

	instancesResumeCommand = instancesCommand.Command("resume", "Resume Instance.")
	instancesResumeName    = instancesResumeCommand.Arg("domain", "Instance domain").Required().String()

	instancesScaleUpCommand = instancesCommand.Command("scaleup", "ScaleUp Instance.")
	instancesScaleUpName    = instancesScaleUpCommand.Arg("name", "Project name.").Required().String()

	instancesScaleDownCommand = instancesCommand.Command("scaledown", "ScaleDown Instance.")
	instancesScaleDownName    = instancesScaleDownCommand.Arg("name", "Project name.").Required().String()
	instancesScaleDownDomain  = instancesScaleDownCommand.Arg("domain", "Domain").Required().String()

	ipsCommand       = app.Command("ips", "IPs.")
	ipsAttachCommand = ipsCommand.Command("attach", "Attach IP to Instance")
	ipsAttachAddress = ipsAttachCommand.Arg("address", "IPV4 Address to attach").Required().String()
	ipsAttachDomain  = ipsAttachCommand.Arg("domain", "Instance domain to attach to").Required().String()

	ipsDetachCommand = ipsCommand.Command("detach", "Detach IP to Instance")
	ipsDetachAddress = ipsDetachCommand.Arg("address", "IPV4 Address to detach").Required().String()

	ipsRequestCommand = ipsCommand.Command("request", "Request an IP")
	ipsReleaseCommand = ipsCommand.Command("release", "Release an IP")
	ipsReleaseAddress = ipsReleaseCommand.Arg("address", "IPV4 Address to release").Required().String()

	ipsListCommand = ipsCommand.Command("list", "List IPs")

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
	volumesDownloadID      = volumesDownloadCommand.Arg("id", "Volume id.").Required().Int()
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
	resourcesNewName    = resourcesNewCommand.Arg("name", "name").Required().String()
	resourcesNewOwner   = resourcesNewCommand.Arg("builtin", "builtin").Required().String()
	resourcesNewBuiltin = resourcesNewCommand.Arg("owner", "owner").Required().String()

	resourcesAvailableCommand = resourcesCommand.Command("available", "List available resources.")

	resourcesListCommand = resourcesCommand.Command("list", "List provisioned resources to project.")
	resourcesListName    = resourcesListCommand.Arg("project_name", "project_name.").String()

	addonsCommand          = app.Command("addons", "Addons.")
	addonsAvailableCommand = addonsCommand.Command("available", "List available addons.")
	addonsListCommand      = addonsCommand.Command("list", "List provisioned addons.")

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

	cli = middleware.NewCLIImplementation(dtoken)
}

var cli *middleware.CLIImplementation

func main() {

	kingpin.Version("0.0.1")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "version":
		fmt.Println("0.0.1")
	case "projects new":
		setToken()
		projects := &Projects{}
		projects.New(*projectsNewName, *projectsNewLanguage, *projectsNewSource, *projectsNewScript)
	case "projects delete":
		setToken()
		projects := &Projects{}
		projects.Delete(*projectsDeleteName)
	case "projects download":
		setToken()
		projects := &Projects{}
		projects.Download(*projectsDownloadName)
	case "projects list":
		setToken()
		projects := &Projects{}
		projects.List()
	case "projects manifest":
		setToken()
		projects := &Projects{}
		projects.Manifest(*projectsManifestName)
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
	case "ips list":
		setToken()
		ips := &Ips{}
		ips.List()
	case "ips request":
		setToken()
		ips := &Ips{}
		ips.Request()
	case "ips release":
		setToken()
		ips := &Ips{}
		ips.Release(*ipsReleaseAddress)
	case "ips attach":
		setToken()
		ips := &Ips{}
		ips.Attach(*ipsAttachAddress, *ipsAttachDomain)
	case "ips detach":
		setToken()
		ips := &Ips{}
		ips.Detach(*ipsDetachAddress)
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
		volumes := &Volumes{}
		volumes.Download(*volumesDownloadID)
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
	case "resources available":
		setToken()
		resources := &Resources{}
		resources.Available()
	case "resources create":
		setToken()
		resources := &Resources{}
		resources.New(*resourcesNewName, *resourcesNewOwner, *resourcesNewBuiltin)
	case "resources list":
		setToken()
		resources := &Resources{}
		if *resourcesListName != "" {
			resources.ListByName(*resourcesListName)
		} else {
			resources.List()
		}
	case "addons available":
		setToken()
		addons := &Addons{}
		addons.Available()
	case "addons list":
		setToken()
		addons := &Addons{}
		addons.List()
	case "status":
		setToken()
		status := &Status{}
		status.Show()
	}
}
