package main

// Instance is the base struct for management of instances
type Instance struct {
	Name   string `json:"Name"`
	Domain string `json:"Domain"`
	Force  bool   `json:"Force"`
}

// Image is the base struct for management of images
type Image struct {
	Name     string `json:"Name"`
	Source   string `json:"Source"`
	Addon    string `json:"Addon"`
	Script   string `json:"Script"`
	Language string `json:"Language"`
	Count    int    `json:"Count"`
	Force    bool   `json:"Force"`
	MakeBin  bool   `json:"MakeBin"`
}

// Resource is the base struct for management of resources
type Resource struct {
	Name    string `json:"Name"`
	Owner   string `json:"Owner"`
	Builtin string `json:"Builtin"`
}
