package codegen

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/Applifier/graphql-codegen/config"
	codegenTemplate "github.com/Applifier/graphql-codegen/template"
	graphql "github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/introspection"
)

type typeConfig struct {
	ignore     bool
	goType     string
	importPath string
}

var (
	internalTypeConfig = map[string]typeConfig{
		"SCALAR": typeConfig{
			true,
			"",
			"",
		},
		"Boolean": typeConfig{
			true,
			"bool",
			"",
		},
		"Float": typeConfig{
			true,
			"float64",
			"",
		},
		"Int": typeConfig{
			true,
			"int32",
			"",
		},
		"ID": typeConfig{
			true,
			"graphql.ID",
			"graphql \"github.com/neelance/graphql-go\"",
		},
		"String": typeConfig{
			true,
			"string",
			"",
		},
	}

	templateFunMap = template.FuncMap{
		"capitalize":         capitalise,
		"uncapitalize":       unCapitalise,
		"is_entry":           isEntryPoint,
		"remove_line_breaks": removeLineBreaks,
	}
)

func Generate(graphSchema string, conf config.Config) (map[string]string, error) {
	sch, err := graphql.ParseSchema(graphSchema, nil)
	if err != nil {
		return nil, err
	}

	ins := sch.Inspect()

	results := map[string]string{}

	var entryPoint = false

	for _, qlType := range ins.Types() {
		name := *qlType.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}

		if internalTypeConfig[name].ignore {
			continue
		}

		if isEntryPoint(name) {
			entryPoint = true
		}

		fileName := fmt.Sprintf("%s_gen.go", strings.ToLower(name))

		log.Printf("Generating Go code for %s %s", qlType.Kind(), name)

		var code string
		code, err = generateType(qlType, conf)
		if err != nil {
			return nil, err
		}

		results[fileName] = code
	}

	// Generate entry point
	if entryPoint {
		entry, err := generateEntryPoint(conf)
		if err != nil {
			return nil, err
		}
		results["resolver_gen.go"] = entry
	}

	return results, nil
}

func returnString(strPtr *string) string {
	if strPtr != nil {
		return *strPtr
	}
	return ""
}

func generateEntryPoint(conf config.Config) (string, error) {
	// TODO figure out if this needs to be configurable
	typeTemplate, err := codegenTemplate.GetTypeTemplate("default")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("default").Funcs(templateFunMap).Parse(strings.Trim(typeTemplate.TypeTemplate, " \t"))
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	tmpl.Execute(buf, map[string]interface{}{
		"Kind":            "RESOLVER",
		"TypeName":        "Resolver",
		"TypeDescription": "Resolver is the main resolver for all queries",
		"Config":          conf,
	})

	b, err := FormatCode(string(buf.Bytes()))
	return string(b), err
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

		tmpl, err := template.New(templateName).Funcs(templateFunMap).Parse(strings.Trim(typeTemplate.TypeTemplate, " \t"))
		if err != nil {
			return "", err
		}

		// Move this to a util func
		var ifields []*introspection.Field
		if tp.Fields(&struct{ IncludeDeprecated bool }{true}) != nil {
			ifields = *tp.Fields(&struct{ IncludeDeprecated bool }{true})
		}

		fields := make([]string, len(ifields))
		methods := make([]string, len(ifields))
		imports := []string{}
		for i, fp := range ifields {
			fieldCode, methodCode, fieldImports, err := generateField(fp, tp, typeConf, conf)
			if err != nil {
				return "", err
			}
			fields[i] = fieldCode
			methods[i] = methodCode

			imports = append(imports, fieldImports...)
		}

		var inputFields []string
		if tp.InputFields() != nil {
			for _, ip := range *tp.InputFields() {
				inputField, inputFieldImports, err := generateInputValue(ip, tp, typeConf, conf)
				if err != nil {
					return "", err
				}
				imports = append(imports, inputFieldImports...)
				inputFields = append(inputFields, inputField)
			}
		}

		possibleTypes := []string{}

		if tp.PossibleTypes() != nil {
			for _, tp := range *tp.PossibleTypes() {
				possibleTypes = append(possibleTypes, *tp.Name())
			}
		}

		enumValues := []string{}
		if tp.EnumValues(&struct{ IncludeDeprecated bool }{true}) != nil {
			for _, value := range *tp.EnumValues(&struct{ IncludeDeprecated bool }{true}) {
				enumValues = append(enumValues, value.Name())
			}
		}

		tmpl.Execute(buf, map[string]interface{}{
			"Kind":            tp.Kind(),
			"PossibleTypes":   possibleTypes,
			"EnumValues":      enumValues,
			"TypeName":        name,
			"TypeDescription": removeLineBreaks(returnString(tp.Description())),
			"Config":          conf,
			"Fields":          fields,
			"InputFields":     inputFields,
			"Methods":         methods,
			"Imports":         removeDuplicates(imports),
		})
	}
	//println(string(buf.Bytes()))
	b, err := FormatCode(string(buf.Bytes()))
	return string(b), err
}

func generateInputValue(ip *introspection.InputValue, tp *introspection.Type, typeConf config.TypeConfig, conf config.Config) (string, []string, error) {
	name := ip.Name()
	propConf := typeConf.Field[name]
	fieldCode := &bytes.Buffer{}
	imports := []string{}

	if len(propConf.Template) == 0 {
		propConf.Template = map[string]map[string]interface{}{}
		propConf.Template["default"] = map[string]interface{}{}
	}

	for templateName, templateConfig := range propConf.Template {
		propTemplate, err := codegenTemplate.GetPropertyTemplate(templateName)
		if err != nil {
			return "", nil, err
		}

		tmpl, err := template.New(templateName).Funcs(templateFunMap).Parse(strings.Trim(propTemplate.FieldTemplate, " \t"))
		if err != nil {
			return "", nil, err
		}

		fieldTypeName := getTypeName(ip.Type(), conf, false)

		tmpl.Execute(fieldCode, map[string]interface{}{
			"TypeKind":         tp.Kind(),
			"FieldName":        name,
			"FieldDescription": removeLineBreaks(returnString(ip.Description())),
			"FieldType":        fieldTypeName,
			"Config":           conf,
			"TemplateConfig":   templateConfig,
		})

		imports = append(imports, getImports(ip.Type(), conf)...)
	}

	return string(fieldCode.Bytes()), imports, nil
}

func generateField(fp *introspection.Field, tp *introspection.Type, typeConf config.TypeConfig, conf config.Config) (string, string, []string, error) {
	name := fp.Name()
	typeName := *tp.Name()
	propConf := typeConf.Field[name]
	fieldCode := &bytes.Buffer{}
	methodCode := &bytes.Buffer{}
	imports := []string{}

	if len(propConf.Template) == 0 {
		propConf.Template = map[string]map[string]interface{}{}
		propConf.Template["default"] = map[string]interface{}{}
	}

	for templateName, templateConfig := range propConf.Template {
		propTemplate, err := codegenTemplate.GetPropertyTemplate(templateName)
		if err != nil {
			return "", "", nil, err
		}

		tmpl, err := template.New(templateName).Funcs(templateFunMap).Parse(strings.Trim(propTemplate.FieldTemplate, " \t"))
		if err != nil {
			return "", "", nil, err
		}

		fieldTypeName := getTypeName(fp.Type(), conf, false)

		type fieldArgument struct {
			Name string
			Type string
		}
		fieldArguments := make([]fieldArgument, 0, len(fp.Args()))

		for _, field := range fp.Args() {
			fieldArguments = append(fieldArguments, fieldArgument{
				Name: field.Name(),
				Type: getTypeName(field.Type(), conf, true),
			})
			imports = append(imports, getImports(field.Type(), conf)...)
		}

		tmpl.Execute(fieldCode, map[string]interface{}{
			"TypeKind":         tp.Kind(),
			"FieldName":        name,
			"FieldDescription": removeLineBreaks(returnString(fp.Description())),
			"FieldType":        fieldTypeName,
			"Config":           conf,
			"TemplateConfig":   templateConfig,
		})

		tmpl, err = template.New(templateName).Funcs(templateFunMap).Parse(propTemplate.MethodTemplate)
		if err != nil {
			return "", "", nil, err
		}

		tmpl.Execute(methodCode, map[string]interface{}{
			"TypeKind":          tp.Kind(),
			"TypeName":          typeName,
			"MethodArguments":   fieldArguments,
			"MethodDescription": removeLineBreaks(returnString(fp.Description())),
			"MethodName":        name,
			"MethodReturnType":  fieldTypeName,
			"MethodReturn":      name,
			"Config":            conf,
			"TemplateConfig":    templateConfig,
		})

		imports = append(imports, getImports(fp.Type(), conf)...)
	}

	return string(fieldCode.Bytes()), string(methodCode.Bytes()), imports, nil
}

func getPointer(typeName string, fp *introspection.Field) string {
	if fp.Type().Kind() == "NON_NULL" {
		return typeName
	}
	return "*" + typeName
}

func getImports(tp *introspection.Type, conf config.Config) []string {
	name := tp.Name()
	if name != nil {
		if val, ok := internalTypeConfig[*name]; ok {
			return []string{val.importPath}
		}
	}

	if tp.OfType() != nil {
		return getImports(tp.OfType(), conf)
	}

	return []string{}
}

func getTypeName(tp *introspection.Type, conf config.Config, input bool) (typ string) {
check:
	if tp.Kind() == "NON_NULL" {
		tp = tp.OfType()
	} else {
		typ = typ + "*"
	}

	if tp.Kind() == "LIST" {
		tp = tp.OfType()
		typ = typ + "[]"
		goto check
	}

	name := tp.Name()
	if val, ok := internalTypeConfig[*name]; ok {
		return typ + val.goType
	}

	if tp.Kind() == "ENUM" {
		typ = typ + *name
	} else if tp.Kind() != "INPUT_OBJECT" {
		typ = typ + *name + "Resolver"
	} else {
		typ = typ + *name
	}

	if typ[0] != '*' && tp.Kind() == "INPUT_OBJECT" {
		typ = "*" + typ
	}

	return
}

func removeDuplicates(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func isEntryPoint(a string) bool {
	return a == "Mutation" || a == "Query"
}

func removeLineBreaks(a string) string {
	return strings.Replace(a, "\n", " ", -1)
}

func capitalise(str string) string {
	if strings.ToLower(str) == "id" {
		return "ID"
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}

func unCapitalise(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}
