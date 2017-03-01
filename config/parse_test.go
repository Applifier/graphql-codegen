package config

import (
	"reflect"
	"testing"
)

func TestConfigParse(t *testing.T) {
	tests := []struct {
		conf     string
		expected Config
	}{
		{
			conf: `
        package = "main"
      `,
			expected: Config{
				Package: "main",
			},
		},
		{
			conf: `
        package = "main"

        type "SomeType" {
          field "someField" {
            template "some_template" {
              arg1 = "arg1"
            }
          }
        }

        type "SomeOtherType" {
          field "someOtherField" {
            template "some_template" {
              arg1 = "arg1"
            }
          }

          field "someOtherField2" {
            template "some_template" {
              arg1 = "arg1"
            }
          }
        }
      `,
			expected: Config{
				Package: "main",
				Type: map[string]TypeConfig{
					"SomeType": TypeConfig{
						Field: map[string]FieldConfig{
							"someField": {
								Template: map[string]map[string]interface{}{
									"some_template": map[string]interface{}{
										"arg1": "arg1",
									},
								},
							},
						},
					},
					"SomeOtherType": TypeConfig{
						Field: map[string]FieldConfig{
							"someOtherField": {
								Template: map[string]map[string]interface{}{
									"some_template": map[string]interface{}{
										"arg1": "arg1",
									},
								},
							},
							"someOtherField2": {
								Template: map[string]map[string]interface{}{
									"some_template": map[string]interface{}{
										"arg1": "arg1",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		cfg, err := Parse(test.conf)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(cfg, test.expected) {
			t.Errorf("Expected\n%+v\nto equal\n%+v", cfg, test.expected)
		}
	}
}
