package codegen

import (
	"github.com/Applifier/graphql-codegen/config"
	"github.com/davecgh/go-spew/spew"
	graphql "github.com/neelance/graphql-go"
)

func Generate(graphSchema string, conf config.Config) (map[string]string, error) {
	sch, err := graphql.ParseSchema(graphSchema, nil)
	if err != nil {
		return nil, err
	}

	ins := sch.Inspect()

	// This is here just for debug
	spew.Dump(ins)

	return nil, nil
}
