package main

import (
	"strconv"
)

// formatSz formats bytes to human readable format
func formatSz(sz int64) string {
	if sz >= 1073741824 {
		return strconv.Itoa(int(sz/1073741824)) + "Gb"
	} else if sz >= 1048576 {
		return strconv.Itoa(int(sz/1048576)) + "Mb"
	} else if sz >= 1024 {
		return strconv.Itoa(int(sz/1024)) + "Kb"
	} else {
		return strconv.Itoa(int(sz))
	}
}
