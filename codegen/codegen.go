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
	goType string
}

var (
	internalTypeConfig = map[string]typeConfig{
		"SCALAR": typeConfig{
			true,
			"",
		},
		"Boolean": typeConfig{
			true,
			"bool",
		},
		"Float": typeConfig{
			true,
			"float64",
		},
		"Int": typeConfig{
			true,
			"int64",
		},
		"ID": typeConfig{
			true,
			"graphql.ID",
		},
		"String": typeConfig{
			true,
			"string",
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

		code, err := generateType(qlType, conf)
		if err != nil {
			return nil, err
		}

		results[fileName] = code
	}

	return results, nil
}

func generateType(tp *introspection.Type, conf config.Config) (code string, err error) {
	name := *tp.Name()
	typeConf := conf.Type[name]

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

		tmpl, err := template.New(templateName).Parse(strings.Trim(typeTemplate.TypeTemplate, " \t"))
		if err != nil {
			return "", err
		}

		// Move this to a util func
		typeName := strings.ToLower(string(name[0])) + name[1:]
		ifields := *tp.Fields(&struct{ IncludeDeprecated bool }{true})
		fields := make([]string, len(ifields))
		methods := make([]string, len(ifields))
		for i, fp := range ifields {
			fieldCode, methodCode, err := generateField(fp, tp, typeConf, conf)
			if err != nil {
				return "", err
			}
			fields[i] = fieldCode
			methods[i] = methodCode
		}

		tmpl.Execute(buf, map[string]interface{}{
			"TypeName":        typeName,
			"TypeDescription": *tp.Description(),
			"Config":          conf,
			"Fields":          fields,
			"Methods":         methods,
		})
	}

	b, err := FormatCode(string(buf.Bytes()))
	return string(b), err
}

func generateField(fp *introspection.Field, tp *introspection.Type, typeConf config.TypeConfig, conf config.Config) (string, string, error) {
	name := fp.Name()
	typeName := *tp.Name()
	propConf := typeConf.Field[name]
	fieldCode := &bytes.Buffer{}
	methodCode := &bytes.Buffer{}

	if len(propConf.Template) == 0 {
		propConf.Template = map[string]map[string]interface{}{}
		propConf.Template["default"] = map[string]interface{}{}
	}

	for templateName, _ := range propConf.Template {
		propTemplate, err := codegenTemplate.GetPropertyTemplate(templateName)
		if err != nil {
			return "", "", err
		}

		tmpl, err := template.New(templateName).Parse(strings.Trim(propTemplate.FieldTemplate, " \t"))
		if err != nil {
			return "", "", err
		}

		// Move this to a util func
		fieldName := strings.ToLower(string(name[0])) + name[1:]

		tmpl.Execute(fieldCode, map[string]interface{}{
			"FieldName":        fieldName,
			"FieldDescription": fp.Description(),
			"FieldType":        getPointer(getTypeName(fp.Type(), conf), fp),
			"Config":           conf,
		})

		tmpl, err = template.New(templateName).Parse(propTemplate.MethodTemplate)
		if err != nil {
			return "", "", err
		}

		tmpl.Execute(methodCode, map[string]interface{}{
			"TypeName":         strings.ToLower(string(typeName[0])) + typeName[1:],
			"MethodName":       strings.ToUpper(string(name[0])) + name[1:],
			"MethodReturnType": getPointer(getTypeName(fp.Type(), conf), fp),
			"MethodReturn":     fieldName,
			"Config":           conf,
		})
	}

	return string(fieldCode.Bytes()), string(methodCode.Bytes()), nil
}

func getPointer(typeName string, fp *introspection.Field) string {
	if fp.Type().Kind() == "NON_NULL" {
		return typeName
	}
	return "*" + typeName
}

func getTypeName(tp *introspection.Type, conf config.Config) string {
	name := tp.Name()
	if name != nil {
		if val, ok := internalTypeConfig[*name]; ok {
			return val.goType
		}
	}

	if tp.OfType() != nil {
		return getTypeName(tp.OfType(), conf)
	}

	return ""
}
