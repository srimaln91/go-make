// Package build provides details of the built binary
// The details are set using ldflags.
//
// The ldflags can be set manually for testing locally:
// `go build -ldflags "-X github.com/srimaln91/go-make.version=$(git describe --tags) -X github.com/srimaln91/go-make.date=$(date -u +%Y-%m-%d-%H:%M:%S-%Z)"`
package build

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/olekukonko/tablewriter"
)

// Details represents known data for a given build
type Details struct {
	Version   string `json:"version,omitempty"`
	GoRuntime string `json:"go_runtime,omitempty"`
	GitCommit string `json:"git_commit,omitempty"`
	OSArch    string `json:"os_arch,omitempty"`
	Date      string `json:"date,omitempty"`
}

type Printer interface {
	Print(details Details) error
}

var version, date, gitCommit, osArch string

var (
	PRINTER_TABLE       = TablePrinter{Writer: os.Stdout}
	PRINTER_STRING      = StringPrinter{Writer: os.Stdout}
	PRINTER_SINGLE_LINE = SingleLinePrinter{Writer: os.Stdout}
)

/*
String returns build details as a string with formatting
suitable for console output.
Ex:
Build Details:
  Version: v0.5.0
  Go Runtime: go1.12.9
  Git Commit: bc2e7ce8edc4aa85cc258890e0e4381630cbf5f8
  OS/Arch: linux/amd64
  Built On: 2019-10-05-12:17:29-UTC
*/
func String() string {
	return fmt.Sprintf(`
Build Details:
  Version: %s
  Go Runtime: %s
  Git Commit: %s
  OS/Arch: %s
  Built on: %s
`,
		version,
		runtime.Version(),
		gitCommit,
		osArch,
		date,
	)
}

/*
Table returns build details as a table
Suitable for console output
Ex:

+----------------+------------+------------------------------------------+-------------+-------------------------+
| BINARY VERSION | GO VERSION |                GIT COMMIT                |   OS/ARCH   |          BUILT          |
+----------------+------------+------------------------------------------+-------------+-------------------------+
| v0.5.0-dirty   | go1.12.9   | bc2e7ce8edc4aa85cc258890e0e4381630cbf5f8 | linux/amd64 | 2019-10-05-12:17:29-UTC |
+----------------+------------+------------------------------------------+-------------+-------------------------+
*/
func Table() string {
	tableBuffer := new(bytes.Buffer)

	data := []string{
		version,
		runtime.Version(),
		gitCommit,
		osArch,
		date,
	}

	table := tablewriter.NewWriter(tableBuffer)

	table.SetHeader([]string{"Binary Version", "Go Version", "Git Commit", "OS/Arch", "Built"})
	table.Append(data)
	table.Render()

	return tableBuffer.String()
}

// JSON returns build details as a JSON string
func JSON() ([]byte, error) {
	return json.Marshal(Data())
}

// Data returns build details as a struct
func Data() Details {
	return Details{
		Version:   version,
		GoRuntime: runtime.Version(),
		GitCommit: gitCommit,
		OSArch:    osArch,
		Date:      date,
	}
}

// CheckVersion checks --version os argument and prints the binary build details in the console
func CheckVersion(printer ...Printer) {

	if len(printer) == 0 {
		printer = append(printer, StringPrinter{
			Writer: os.Stdout,
		})
	}

	// Check OS arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--version", "-v":
			for _, p := range printer {
				err := p.Print(Data())
				if err != nil {
					os.Exit(1)
				}
			}
			os.Exit(0)
		}
	}
}
