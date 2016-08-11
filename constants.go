package main

var (
	// APIVersion is the version of this CLI
	APIVersion = "v0.1"

	// APIBase is the base url that cli requests goto
	//	APIBase = "https://deferpanic.net/" + APIVersion
	APIBase = "http://127.0.0.1:3000/" + APIVersion

	// imageURL is the url for image management
	imageURL = APIBase + "/image"

	// instanceURL is the url for instance management
	instanceURL = APIBase + "/instance"

	// URL is the url for backup management
	backupURL = APIBase + "/storage"

	// systemURL is the url for system management
	systemURL = APIBase + "/system"

	// getURL is the url to download rumprun image binary
	getURL = imageURL + "/get"

	// putURL is the url to upload rumprun image binary
	putURL = imageURL + "/put"

	// languagesURL is the url to view system languages
	languagesURL = systemURL + "/languages"

	// addonsURL is the url to view system addons
	addonsURL = systemURL + "/addons"

	// builtinsURL is the url to view system builtins
	builtinsURL = systemURL + "/builtins"

	// statusURL is the url to view status of the system
	statusURL = systemURL + "/status"
)
