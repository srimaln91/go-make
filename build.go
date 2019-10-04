// Package build provides details of the built binary
// The details are set using ldflags.
//
// The ldflags can be set manually for testing locally:
// `go build -ldflags "-X github.com/srimaln91/go-build/util/build.version=$(git describe --tags) -X github.com/srimaln91/go-build/util/build.date=$(date -u +%Y-%m-%d-%H:%M:%S-%Z)"`
package build

import (
	"bytes"
	"fmt"
	"os"
	"runtime"

	"github.com/kataras/tablewriter"
)

// Details represents known data for a given build
type Details struct {
	Version   string `json:"version,omitempty"`
	GoVersion string `json:"go_version,omitempty"`
	GitCommit string `json:"git_commit,omitempty"`
	OSArch    string `json:"os_arch,omitempty"`
	Date      string `json:"date,omitempty"`
}

var goVersion = runtime.Version()
var version, date, gitCommit, osArch string

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
  Go Version: %s
  Git Commit: %s
  OS/Arch: %s
  Build Date: %s
`,
		version,
		goVersion,
		gitCommit,
		osArch,
		date,
	)
}

// Table returns build details as a styled table
func Table() string {
	tableBuffer := new(bytes.Buffer)

	data := []string{version, goVersion, gitCommit, osArch, date}
	table := tablewriter.NewWriter(tableBuffer)

	table.SetHeader([]string{"Binary Version", "Go Version", "Git Commit", "OS/Arch", "Build Date"})
	table.Append(data)
	table.Render()

	return tableBuffer.String()
}

// Data returns build details as a struct
func Data() Details {
	return Details{
		Version:   version,
		GoVersion: goVersion,
		GitCommit: gitCommit,
		OSArch:    osArch,
		Date:      date,
	}
}

// CheckVersion checks --version os argument and prints the binary version
func CheckVersion() {

	// Check OS arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--version":
			fmt.Fprintf(os.Stdout, "%s\n", Table())
			os.Exit(0)
		}
	}
}
