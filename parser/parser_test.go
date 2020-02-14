package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func parseFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if password := Parse(content); password != "F3RN3T4NDC0C4C0L4" {
			return fmt.Errorf("expected password to be F3RN3T4NDC0C4C0L4 but got %s", password)
		}
	}
	return nil
}

func TestParsingPassword(t *testing.T) {
	if err := filepath.Walk("./__data__", parseFile); err != nil {
		t.Fatal(err)
	}
}
