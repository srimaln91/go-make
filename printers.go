package build

import (
	"fmt"
	"io"
)

type TablePrinter struct {
	Writer io.Writer
}

func (p TablePrinter) Print(details Details) error {
	_, err := fmt.Fprintf(p.Writer, "%s\n", Table())
	return err
}

type SingleLinePrinter struct {
	Writer io.Writer
}

func (p SingleLinePrinter) Print(details Details) error {
	_, err := fmt.Fprintf(p.Writer, "%s\n", fmt.Sprintf("Version: %s", details.Version))
	return err
}

type StringPrinter struct {
	Writer io.Writer
}

func (p StringPrinter) Print(details Details) error {
	_, err := fmt.Fprintf(p.Writer, "%s\n", String())
	return err
}
