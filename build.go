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
	"io"
	"os"
	"runtime"

	"github.com/olekukonko/tablewriter"
)

// Details represents known data for a given build
type Details struct {
	Version   string `json:"version,omitempty"`
	GoVersion string `json:"go_version,omitempty"`
	GitCommit string `json:"git_commit,omitempty"`
	OSArch    string `json:"os_arch,omitempty"`
	Date      string `json:"date,omitempty"`
}

var version, date, gitCommit, osArch string

type Printer interface {
	Print(details Details) error
}

type BasePrinter struct {
	Writer io.Writer
}

type TablePrinter struct {
	BasePrinter
}

func (p TablePrinter) Print(details Details) error {
	_, err := fmt.Fprintf(p.Writer, "%s\n", Table())
	return err
}

type SingleLinePrinter struct {
	BasePrinter
}

func (p SingleLinePrinter) Print(details Details) error {
	_, err := fmt.Fprintf(p.Writer, "%s\n", fmt.Sprintf("Version: %s", details.Version))
	return err
}

type StringPrinter struct {
	BasePrinter
}

func (p StringPrinter) Print(details Details) error {
	_, err := fmt.Fprintf(p.Writer, "%s\n", String())
	return err
}

/*
String returns build details as a string with formatting
suitable for console output.
Ex:
Build Details:
  Version: v0.5.0
  Go Version: go1.12.9
  Git Commit: bc2e7ce8edc4aa85cc258890e0e4381630cbf5f8
  OS/Arch: linux/amd64
  Built: 2019-10-05-12:17:29-UTC
*/
func String() string {
	return fmt.Sprintf(`
Build Details:
  Version: %s
  Go Version: %s
  Git Commit: %s
  OS/Arch: %s
  Built: %s
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
		GoVersion: runtime.Version(),
		GitCommit: gitCommit,
		OSArch:    osArch,
		Date:      date,
	}
}

// CheckVersion checks --version os argument and prints the binary build details in the console
func CheckVersion(printer ...Printer) {

	if len(printer) == 0 {
		printer = append(printer, SingleLinePrinter{
			BasePrinter: BasePrinter{Writer: os.Stdout},
		})
	}

	// Check OS arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--version":
			for _, p := range printer {
				p.Print(Data())
			}
			os.Exit(0)
		}
	}
}
