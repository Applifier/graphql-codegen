package codegen

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Applifier/graphql-codegen/config"
	codegenTemplate "github.com/Applifier/graphql-codegen/template"
	graphql "github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/introspection"
)

type typeConfig struct {
	ignore bool
}

var (
	internalTypeConfig = map[string]typeConfig{
		"SCALAR": typeConfig{
			true,
		},
		"Boolean": typeConfig{
			true,
		},
		"Float": typeConfig{
			true,
		},
		"Int": typeConfig{
			true,
		},
		"ID": typeConfig{
			true,
		},
		"String": typeConfig{
			true,
		},
	}
)

func Generate(graphSchema string, conf config.Config) (map[string]string, error) {
	sch, err := graphql.ParseSchema(graphSchema, nil)
	if err != nil {
		return nil, err
	}

	ins := sch.Inspect()

	results := map[string]string{}

	for _, qlType := range ins.Types() {
		name := *qlType.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}

		if internalTypeConfig[name].ignore {
			continue
		}

		fileName := fmt.Sprintf("%s_gen.go", strings.ToLower(name))

		code, err := generateType(qlType, conf.Type[name])
		if err != nil {
			return nil, err
		}

		results[fileName] = code
	}

	return results, nil
}

func generateType(tp *introspection.Type, typeConf config.TypeConfig) (code string, err error) {
	buf := &bytes.Buffer{}

	if len(typeConf.Template) == 0 {
		typeConf.Template = map[string]map[string]interface{}{}
		typeConf.Template["default"] = map[string]interface{}{}
	}

	for templateName, _ := range typeConf.Template {
		typeTemplate, err := codegenTemplate.GetTypeTemplate(templateName)
		if err != nil {
			return "", err
		}

		tmpl, err := template.New(templateName).Parse(typeTemplate.TypeTemplate)
		if err != nil {
			return "", err
		}

		tmpl.Execute(buf, "foo")
	}
	b, err := FormatCode(string(buf.Bytes()))
	return string(b), err
}
