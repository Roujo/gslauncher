package version

import "fmt"

// When updating this version number also update installer.nsi
var Major int = 1
var Minor int = 3
var Patch int = 1

var Protocol int = 1

func Formatted() string {
	return fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
}
