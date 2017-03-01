package config

import "github.com/hashicorp/hcl"

// Parse parses codegen config
func Parse(config string) (cfg Config, err error) {
	err = hcl.Decode(&cfg, config)
	return
}
