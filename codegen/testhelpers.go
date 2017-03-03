package codegen

import (
	"testing"

	"github.com/Applifier/graphql-codegen/config"
)

func RunTest(schema string, conf config.Config, expected map[string]string, t *testing.T) map[string]string {
	fileMap, err := NewCodeGen(schema, conf).Generate()
	if err != nil {
		t.Log("Schema parsing failed")
		t.Fatal(err)
	}

	for file, resultingCode := range fileMap {
		code, err := FormatCode(expected[file])
		if err != nil {
			t.Fatal(err)
		}

		if string(code) != resultingCode {
			t.Errorf("Generated file %s content\n%s\n\nshould have matched\n\n%s", file, resultingCode, code)
		}
	}

	return fileMap
}
