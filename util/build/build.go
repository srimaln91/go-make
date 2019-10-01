// Package build provides details of the built binary
// The details are set using ldflags.
//
// The ldflags can be set manually for testing locally:
// `go build -ldflags "-X git.mytaxi.lk/pickme/geo/services/shuttle-location-handler/util/build.version=$(git describe --tags) -X git.mytaxi.lk/pickme/geo/services/shuttle-location-handler/util/build.date=$(date -u +%Y-%m-%d-%H:%M:%S-%Z)"`
package build

import (
	"fmt"
)

// Details represents known data for a given build
type Details struct {
	Version     string `json:"version,omitempty"`
	Date        string `json:"date,omitempty"`
	Name        string `json:"name,omitempty"`
	Discription string `json:"discription,omitempty"`
}

var version, date string

const name = "Put binary name here"
const discription = "Put description here"

// String returns build details as a string with formatting
// suitable for console output.
//
// i.e.
// Build Details:
//         Version:        v0.1.0-155-g1a20f8b
//         Date:           2018-11-05-14:33:14-UTC
func String() string {
	return fmt.Sprintf(`
Build Details:
  Version: %s
  Build Date: %s
`,
		version,
		date,
	)
}

// Data returns build details as a struct
func Data() Details {
	return Details{
		Version:     version,
		Date:        date,
		Name:        name,
		Discription: discription,
	}
}
