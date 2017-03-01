package codegen

import (
	"github.com/davecgh/go-spew/spew"
	graphql "github.com/neelance/graphql-go"
)

func Generate(graphSchema string, config Config) (map[string]string, error) {
	sch, err := graphql.ParseSchema(graphSchema, nil)
	if err != nil {
		return nil, err
	}

	ins := sch.Inspect()

	// This is here just for debug
	spew.Dump(ins)

	return nil, nil
}
