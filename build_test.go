package build

import (
	"fmt"
	"testing"
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
