package build

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"
	"time"
)

func TestTable(t *testing.T) {
	tableString := Table()

	if tableString == "" {
		t.Error("Error in Table()")
	}

	fmt.Println(tableString)
}

func TestString(t *testing.T) {
	resultString := String()
	if resultString == "" {
		t.Error("Error in String()")
	}

	fmt.Println(resultString)
}

func TestJSON(t *testing.T) {
	resultJSON, err := JSON()
	if err != nil {
		t.Error("Error in JSON()", err)
	}

	if string(resultJSON) == "" {
		t.Error("Error in JSON()", "Empty string received")
	}

	fmt.Println(string(resultJSON))
}

func TestSingleLinePrinter(t *testing.T) {
	d := Details{
		Version:   "v1.0.0",
		GoVersion: runtime.Version(),
		GitCommit: "test got commit",
		OSArch:    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Date:      time.Now().String(),
	}

	writer := bytes.NewBuffer(nil)
	p := SingleLinePrinter{Writer: writer}
	p.Print(d)

	b, err := ioutil.ReadAll(writer)
	if err != nil {
		t.Error(err)
	}
	if b == nil {
		t.Error("empty output from printer")
	}
}

func TestTablePrinter(t *testing.T) {
	d := Details{
		Version:   "v1.0.0",
		GoVersion: runtime.Version(),
		GitCommit: "test got commit",
		OSArch:    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Date:      time.Now().String(),
	}

	writer := bytes.NewBuffer(nil)
	p := TablePrinter{Writer: writer}
	p.Print(d)

	b, err := ioutil.ReadAll(writer)
	if err != nil {
		t.Error(err)
	}
	if b == nil {
		t.Error("empty output from printer")
	}
}

func TestStringPrinter(t *testing.T) {
	d := Details{
		Version:   "v1.0.0",
		GoVersion: runtime.Version(),
		GitCommit: "test got commit",
		OSArch:    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Date:      time.Now().String(),
	}

	writer := bytes.NewBuffer(nil)
	p := StringPrinter{Writer: writer}
	p.Print(d)

	b, err := ioutil.ReadAll(writer)
	if err != nil {
		t.Error(err)
	}
	if b == nil {
		t.Error("empty output from printer")
	}
}
