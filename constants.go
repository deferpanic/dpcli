package main

var (
	// APIVersion is the version of this CLI
	APIVersion = "v0.1"

	// APIBase is the base url that cli requests goto
	APIBase = "http://deferpanic.net/" + APIVersion

	// imageURL is the url for image management
	imageURL = APIBase + "/image"

	// instanceURL is the url for instance management
	instanceURL = APIBase + "/instance"

	// URL is the url for backup management
	backupURL = APIBase + "/storage"

	// systemURL is the url for system management
	systemURL = APIBase + "/system"

	// resourceURL is the url for resource management
	resourceURL = APIBase + "/resource"

	// getURL is the url to download rumprun image binary
	getURL = imageURL + "/get"

	// putURL is the url to upload rumprun image binary
	putURL = imageURL + "/put"

	// scaleupURL is the url to scale up execution of rumprun image
	scaleupURL = instanceURL + "/scaleup"

	// scaledownURL is the url to scale down execution of rumprun image
	scaledownURL = instanceURL + "/scaledown"

	// runlogURL is the url to view execution log of rumprun image
	runlogURL = instanceURL + "/log"

	// showURL is the url to show running rumprun images
	showURL = instanceURL + "/show"

	// pauseURL is the url to stop existing rumprun image
	pauseURL = instanceURL + "/pause"

	// resumeURL is the url to start existing rumprun image
	resumeURL = instanceURL + "/resume"

	// savebackupURL is the url to save a backup for running rumprun image
	savebackupURL = backupURL + "/save"

	// restorebackupURL is the url to restore a backup for running rumprun image
	restorebackupURL = backupURL + "/restore"

	// listbackupsURL is the url to list all rumprun image backups
	listbackupsURL = backupURL + "/list"

	// newresourceURL is the url to add a resource for rumprun image
	newresourceURL = resourceURL + "/new"

	// listresourcesURL is the url to list all rumprun image resources
	listresourcesURL = resourceURL + "/list"

	// languagesURL is the url to view system languages
	languagesURL = systemURL + "/languages"

	// addonsURL is the url to view system addons
	addonsURL = systemURL + "/addons"

	// builtinsURL is the url to view system builtins
	builtinsURL = systemURL + "/builtins"

	// statusURL is the url to view status of the system
	statusURL = systemURL + "/status"

	// URL is the url for volume management
	volumeURL = APIBase + "/volume"

	// newvolumeURL is the url to add a volume for rumprun image
	newvolumeURL = volumeURL + "/new"

	// removevolumeURL is the url to remove a volume from rumprun image
	removevolumeURL = volumeURL + "/remove"

	// connectvolumeURL is the url to connect a volume to rumprun image
	connectvolumeURL = volumeURL + "/connect"

	// disconnectvolumeURL is the url to disconnect a volume from
	// rumprun image
	disconnectvolumeURL = volumeURL + "/disconnect"

	// listvolumesURL is the url to list all rumprun image volumes
	listvolumesURL = volumeURL + "/list"

	// getvolumeURL is the url to download rumprun image volume
	getvolumeURL = volumeURL + "/get"

	// putvolumeURL is the url to upload rumprun image volume
	putvolumeURL = volumeURL + "/put"
)
