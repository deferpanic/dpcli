package api

// Image is the base struct for management of images
type Image struct {
	Name          string `json:"Name"`
	Source        string `json:"Source"`
	Addon         string `json:"Addon"`
	Script        string `json:"Script"`
	Language      string `json:"Language"`
	Count         int    `json:"Count"`
	Force         bool   `json:"Force"`
	MakeBin       bool   `json:"MakeBin"`
	SystemVolumes bool   `json:"SystemVolumes"`
	Buildable     bool   `json:"Buildable"`
	Filename      string `json:"Filename"`
}
