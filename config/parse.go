package config

import (
	"fmt"

	"github.com/hashicorp/hcl"
)

// Parse parses codegen config
func Parse(config string) (cfg Config, err error) {
	err = hcl.Decode(&cfg, config)

	broad := map[string]interface{}{}

	hcl.Decode(&broad, config)

	fmt.Printf("%+v\n", broad)

	return
}
