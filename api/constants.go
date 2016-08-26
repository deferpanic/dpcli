package api

var (
	// APIVersion is the version of this Cli
	APIVersion = "v0.1"

	// APIBase is the base url that Cli requests goto
	APIBase = "https://deferpanic.net/" + APIVersion

	// imageURL is the url for image management
	imageURL = APIBase + "/image"

	// instanceURL is the url for instance management
	instanceURL = APIBase + "/instance"

	// URL is the url for backup management
	backupURL = APIBase + "/storage"

	// systemURL is the url for system management
	systemURL = APIBase + "/system"

	// putURL is the url to upload rumprun image binary
	putURL = imageURL + "/put"

	// languagesURL is the url to view system languages
	languagesURL = systemURL + "/languages"

	// builtinsURL is the url to view system builtins
	builtinsURL = systemURL + "/builtins"

	// statusURL is the url to view status of the system
	statusURL = systemURL + "/status"
)
