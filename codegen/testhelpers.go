package codegen

import "testing"

func RunTest(schema string, config Config, expected map[string]string, t *testing.T) {
	fileMap, err := Generate(schema, config)
	if err != nil {
		t.Fatal(err)
	}

	for file, expectedCode := range expected {
		code, err := FormatCode(expectedCode)
		if err != nil {
			t.Fatal(err)
		}

		if string(code) != fileMap[file] {
			t.Errorf("Generated file %s content\n%s\n\nshould have matched\n\n%s", file, fileMap[file], code)
		}
	}
}
