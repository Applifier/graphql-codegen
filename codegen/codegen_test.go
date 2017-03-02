package codegen

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/Applifier/graphql-codegen/config"
)

func TestCodegen(t *testing.T) {
	testDirs, err := ioutil.ReadDir("_fixtures")
	if err != nil {
		t.Fatal(err)
	}
	for _, testDir := range testDirs {
		if strings.HasPrefix(testDir.Name(), "_") && os.Getenv("UNSTABLE_TEST") != "yes" {
			continue
		}

		schemaBytes, _ := ioutil.ReadFile(path.Join("_fixtures", testDir.Name(), "schema.graphql"))
		confBytes, _ := ioutil.ReadFile(path.Join("_fixtures", testDir.Name(), "config.hcl"))
		exectedFiles, _ := ioutil.ReadDir(path.Join("_fixtures", testDir.Name()))
		exectedFilesMap := map[string]string{}

		for _, f := range exectedFiles {
			if strings.HasSuffix(f.Name(), ".go") {
				expectedBytes, _ := ioutil.ReadFile(path.Join("_fixtures", testDir.Name(), f.Name()))
				exectedFilesMap[f.Name()] = string(expectedBytes)
			}
		}

		cfg, _ := config.Parse(string(confBytes))

		fileMap := RunTest(string(schemaBytes), cfg, exectedFilesMap, t)
		if os.Getenv("RECORD_FIXTURES") == "yes" {
			for filename, data := range fileMap {
				ioutil.WriteFile(path.Join("_fixtures", testDir.Name(), filename), []byte(data), 0644)
			}
		}
	}
}
