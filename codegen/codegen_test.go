package codegen

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/Applifier/graphql-codegen/config"
)

func TestCodegen(t *testing.T) {
	testDirs, err := ioutil.ReadDir("fixtures")
	if err != nil {
		t.Fatal(err)
	}
	for _, testDir := range testDirs {
		schemaBytes, _ := ioutil.ReadFile(path.Join("fixtures", testDir.Name(), "schema.graphql"))
		confBytes, _ := ioutil.ReadFile(path.Join("fixtures", testDir.Name(), "config.hcl"))
		exectedFiles, _ := ioutil.ReadDir(path.Join("fixtures", testDir.Name()))
		exectedFilesMap := map[string]string{}

		for _, f := range exectedFiles {
			if strings.HasSuffix(f.Name(), ".go") {
				expectedBytes, _ := ioutil.ReadFile(path.Join("fixtures", testDir.Name(), f.Name()))
				exectedFilesMap[f.Name()] = string(expectedBytes)
			}
		}

		cfg, _ := config.Parse(string(confBytes))

		RunTest(string(schemaBytes), cfg, exectedFilesMap, t)
	}
}
