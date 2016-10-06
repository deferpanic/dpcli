// Package main implements access to the deferpanic unikernel IaaS API
// for users.
package main

import (
	"fmt"
	"github.com/deferpanic/dpcli/api"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"strings"
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
	projectsNewCompiler = projectsNewCommand.Arg("compiler", "Project compiler.").Required().String()
	projectsNewSource   = projectsNewCommand.Arg("source", "Project source.").Required().String()
	projectsNewScript   = projectsNewCommand.Arg("script", "Project script.").String()

	projectsDeleteCommand = projectsCommand.Command("delete", "Delete a project.")
	projectsDeleteName    = projectsDeleteCommand.Arg("name", "Project name.").Required().String()

	projectsDownloadCommand = projectsCommand.Command("download", "Download image.")
	projectsDownloadName    = projectsDownloadCommand.Arg("name", "Project name.").Required().String()
	projectsUploadCommand   = projectsCommand.Command("upload", "Upload image.")
	projectsUploadBinary    = projectsUploadCommand.Arg("binary", "Image binary path.").Required().String()

	projectsListCommand = projectsCommand.Command("list", "List Projects.")
	projectsLogCommand  = projectsCommand.Command("log", "View Latest Project Build Log")
	projectsLogName     = projectsLogCommand.Arg("name", "Project name.").Required().String()

	projectsManifestCommand = projectsCommand.Command("manifest", "Project manifest.")
	projectsManifestName    = projectsManifestCommand.Arg("name", "Project name.").Required().String()

	usersCommand        = app.Command("users", "Users.")
	usersCreateCommand  = usersCommand.Command("create", "Create a new user.")
	usersCreateEmail    = usersCreateCommand.Arg("email", "Email.").Required().String()
	usersCreateUsername = usersCreateCommand.Arg("username", "Username.").Required().String()
	usersCreatePassword = usersCreateCommand.Arg("password", "Password.").Required().String()

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
	compilersCommand = app.Command("compilers", "Compilers.")

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

	searchCommand      = app.Command("search", "Search for a project")
	searchCommandName  = searchCommand.Arg("description", "Description").Required().String()
	searchCommandStars = searchCommand.Arg("stars", "Star Count").Int()

	status = app.Command("status", "Show Status.")
)

// fixme
// should only have to be called once..
func setToken() {
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.dprc")
	if err != nil {
		fmt.Println(api.RedBold("Have an account yet?\n" +
			"If so you can stick your token in ~/.dprc.\n" +
			"Otherwise signup via:\n\n\tdpcli users create my@email.com username password\n"))

	}
	dtoken := string(dat)

	if dtoken == "" {
		dtoken = *token
	}

	if dtoken == "" {
		api.RedBold("no token")
		os.Exit(1)
	}

	dtoken = strings.TrimSpace(dtoken)
	api.Cli = api.NewCliImplementation(dtoken)
}

func main() {

	kingpin.Version("0.0.1")

	if (len(os.Args) > 1) && (os.Args[1] == "users" && os.Args[2] == "create") {
		api.Cli = api.NewCliImplementation("")
	} else {
		setToken()
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "version":
		fmt.Println("0.0.1")
	case "projects new":
		projects := &api.Projects{}
		projects.New(*projectsNewName, *projectsNewLanguage, *projectsNewCompiler, *projectsNewSource, *projectsNewScript)
	case "projects delete":
		projects := &api.Projects{}
		projects.Delete(*projectsDeleteName)
	case "projects download":
		projects := &api.Projects{}
		projects.Download(*projectsDownloadName, ".")
	case "projects list":
		projects := &api.Projects{}
		projects.List()
	case "projects manifest":
		projects := &api.Projects{}
		projects.Manifest(*projectsManifestName)
	case "projects log":
		projects := &api.Projects{}
		projects.Log(*projectsLogName)
	case "users create":
		users := &api.Users{}
		users.Create(*usersCreateEmail, *usersCreateUsername, *usersCreatePassword)
	case "instances scaleup":
		instances := &api.Instances{}
		instances.ScaleUp(*instancesScaleUpName)
	case "instances scaledown":
		instances := &api.Instances{}
		instances.ScaleDown(*instancesScaleDownName, *instancesScaleDownDomain)
	case "instances new":
		instances := &api.Instances{}
		instances.New(*instancesNewName)
	case "instances log":
		instances := &api.Instances{}
		instances.Log(*instancesLogName)
	case "instances list":
		instances := &api.Instances{}
		instances.List(*instancesListName)
	case "instances pause":
		instances := &api.Instances{}
		instances.Pause(*instancesPauseName)
	case "instances resume":
		instances := &api.Instances{}
		instances.Resume(*instancesResumeName)
	case "ips list":
		ips := &api.Ips{}
		ips.List()
	case "ips request":
		ips := &api.Ips{}
		ips.Request()
	case "ips release":
		ips := &api.Ips{}
		ips.Release(*ipsReleaseAddress)
	case "ips attach":
		ips := &api.Ips{}
		ips.Attach(*ipsAttachAddress, *ipsAttachDomain)
	case "ips detach":
		ips := &api.Ips{}
		ips.Detach(*ipsDetachAddress)
	case "volumes list":
		volumes := &api.Volumes{}
		if *volumesListName != "" {
			volumes.ListByName(*volumesListName)
		}
		if *volumesListDomain != "" {
			volumes.ListByDomain(*volumesListDomain)
		}
	case "volumes create":
	case "volumes show":
	case "volumes update":
	case "volumes delete":
	case "volumes attach":
		volumes := &api.Volumes{}
		volumes.Attach(*volumesAttachName, *volumesAttachDomain)
	case "volumes detach":
		volumes := &api.Volumes{}
		volumes.Detach(*volumesAttachName, *volumesAttachDomain)
	case "volumes download":
		volumes := &api.Volumes{}
		volumes.Download(*volumesDownloadID)
	case "volumes upload":
	case "backups list":
		backups := &api.Backups{}
		backups.List()
	case "backups save":
		backups := &api.Backups{}
		backups.Save(*backupsSaveName, *backupsSaveDomain)
	case "backups restore":
		backups := &api.Backups{}
		backups.Restore(*backupsRestoreName, *backupsRestoreDomain)
	case "languages":
		languages := &api.Languages{}
		languages.List()
	case "compilers":
		compilers := &api.Compilers{}
		compilers.List()
	case "resources available":
		resources := &api.Resources{}
		resources.Available()
	case "resources create":
		resources := &api.Resources{}
		resources.New(*resourcesNewName, *resourcesNewOwner, *resourcesNewBuiltin)
	case "resources list":
		resources := &api.Resources{}
		if *resourcesListName != "" {
			resources.ListByName(*resourcesListName)
		} else {
			resources.List()
		}
	case "addons available":
		addons := &api.Addons{}
		addons.Available()
	case "addons list":
		addons := &api.Addons{}
		addons.List()
	case "search":
		search := &api.Search{}
		if *searchCommandStars != 0 {
			search.FindWithStars(*searchCommandName, *searchCommandStars)
		} else {
			search.Find(*searchCommandName)
		}
	case "status":
		status := &api.Status{}
		status.Show()
	}
}
