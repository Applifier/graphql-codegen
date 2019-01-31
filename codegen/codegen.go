package codegen

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/Applifier/graphql-codegen/config"
	codegenTemplate "github.com/Applifier/graphql-codegen/template"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/introspection"
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
			"graphql \"github.com/graph-gophers/graphql-go\"",
		},
		"String": typeConfig{
			true,
			"string",
			"",
		},
	}
)

type CodeGen struct {
	graphSchema  string
	conf         config.Config
	mutationName string
	queryName    string
}

func NewCodeGen(graphSchema string, conf config.Config) *CodeGen {
	return &CodeGen{graphSchema, conf, "", ""}
}

func (g *CodeGen) Generate() (map[string]string, error) {
	graphSchema := g.graphSchema
	conf := g.conf

	sch, err := graphql.ParseSchema(graphSchema, nil)
	if err != nil {
		return nil, err
	}

	ins := sch.Inspect()

	if ins.MutationType() != nil {
		g.mutationName = g.returnString(ins.MutationType().Name())
	}

	if ins.QueryType() != nil {
		g.queryName = g.returnString(ins.QueryType().Name())
	}

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

		if g.isEntryPoint(name) {
			entryPoint = true
		}

		fileName := fmt.Sprintf("%s_gen.go", strings.ToLower(name))

		log.Printf("Generating Go code for %s %s", qlType.Kind(), name)

		var code string
		code, err = g.generateType(qlType, conf)
		if err != nil {
			return nil, err
		}

		results[fileName] = code
	}

	// Generate entry point
	if entryPoint {
		entry, err := g.generateEntryPoint(conf)
		if err != nil {
			return nil, err
		}
		results["resolver_gen.go"] = entry
	}

	return results, nil
}

func (g *CodeGen) returnString(strPtr *string) string {
	if strPtr != nil {
		return *strPtr
	}
	return ""
}

func (g *CodeGen) generateEntryPoint(conf config.Config) (string, error) {
	// TODO figure out if this needs to be configurable
	typeTemplate, err := codegenTemplate.GetTypeTemplate("default")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("default").Funcs(g.templateFuncMap()).Parse(strings.Trim(typeTemplate.TypeTemplate, " \t"))
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

func (g *CodeGen) generateType(tp *introspection.Type, conf config.Config) (code string, err error) {
	name := *tp.Name()
	typeConf := conf.Type[name]

	buf := &bytes.Buffer{}

	if len(typeConf.Template) == 0 {
		typeConf.Template = map[string]map[string]interface{}{}
		typeConf.Template["default"] = map[string]interface{}{}
	}

	for templateName, templateConfig := range typeConf.Template {
		typeTemplate, err := codegenTemplate.GetTypeTemplate(templateName)
		if err != nil {
			return "", err
		}

		tmpl, err := template.New(templateName).Funcs(g.templateFuncMap()).Parse(strings.Trim(typeTemplate.TypeTemplate, " \t"))
		if err != nil {
			return "", err
		}

		// Move this to a util func (g *CodeGen)
		var ifields []*introspection.Field
		if tp.Fields(&struct{ IncludeDeprecated bool }{true}) != nil {
			ifields = *tp.Fields(&struct{ IncludeDeprecated bool }{true})
		}

		fields := make([]string, len(ifields))
		methods := make([]string, len(ifields))
		imports := []string{}
		for i, fp := range ifields {
			fieldCode, methodCode, fieldImports, err := g.generateField(fp, tp, typeConf, conf)
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
				inputField, inputFieldImports, err := g.generateInputValue(ip, tp, typeConf, conf)
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

		imports = append(imports, typeTemplate.Config.Imports...)
		if val, ok := templateConfig["imports"]; ok {
			if arr, ok := val.([]string); ok {
				imports = append(imports, arr...)
			}
		}

		imports = append(imports, typeConf.Imports...)

		tmpl.Execute(buf, map[string]interface{}{
			"Kind":            tp.Kind(),
			"PossibleTypes":   possibleTypes,
			"EnumValues":      enumValues,
			"TypeName":        name,
			"TypeDescription": g.removeLineBreaks(g.returnString(tp.Description())),
			"Config":          conf,
			"Fields":          fields,
			"InputFields":     inputFields,
			"Methods":         methods,
			"Imports":         g.removeDuplicates(imports),
			"TemplateConfig":  templateConfig,
		})
	}
	//println(string(buf.Bytes()))
	b, err := FormatCode(string(buf.Bytes()))
	return string(b), err
}

func (g *CodeGen) generateInputValue(ip *introspection.InputValue, tp *introspection.Type, typeConf config.TypeConfig, conf config.Config) (string, []string, error) {
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

		tmpl, err := template.New(templateName).Funcs(g.templateFuncMap()).Parse(strings.Trim(propTemplate.FieldTemplate, " \t"))
		if err != nil {
			return "", nil, err
		}

		fieldTypeName := g.getTypeName(ip.Type(), conf, false)

		tmpl.Execute(fieldCode, map[string]interface{}{
			"TypeKind":         tp.Kind(),
			"FieldName":        name,
			"FieldDescription": g.removeLineBreaks(g.returnString(ip.Description())),
			"FieldType":        fieldTypeName,
			"Config":           conf,
			"TemplateConfig":   templateConfig,
		})

		imports = append(imports, g.getImports(ip.Type(), conf)...)
		imports = append(imports, propConf.Imports...)
		imports = append(imports, propTemplate.Config.Imports...)
		if val, ok := templateConfig["imports"]; ok {
			if arr, ok := val.([]string); ok {
				imports = append(imports, arr...)
			}
		}
	}

	return string(fieldCode.Bytes()), imports, nil
}

func (g *CodeGen) generateField(fp *introspection.Field, tp *introspection.Type, typeConf config.TypeConfig, conf config.Config) (string, string, []string, error) {
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

		tmpl, err := template.New(templateName).Funcs(g.templateFuncMap()).Parse(strings.Trim(propTemplate.FieldTemplate, " \t"))
		if err != nil {
			return "", "", nil, err
		}

		fieldTypeName := g.getTypeName(fp.Type(), conf, false)

		type fieldArgument struct {
			Name string
			Type string
		}
		fieldArguments := make([]fieldArgument, 0, len(fp.Args()))

		for _, field := range fp.Args() {
			fieldArguments = append(fieldArguments, fieldArgument{
				Name: field.Name(),
				Type: g.getTypeName(field.Type(), conf, true),
			})
			imports = append(imports, g.getImports(field.Type(), conf)...)
		}

		tmpl.Execute(fieldCode, map[string]interface{}{
			"TypeKind":         tp.Kind(),
			"FieldName":        name,
			"FieldDescription": g.removeLineBreaks(g.returnString(fp.Description())),
			"FieldType":        fieldTypeName,
			"Config":           conf,
			"TemplateConfig":   templateConfig,
		})

		tmpl, err = template.New(templateName).Funcs(g.templateFuncMap()).Parse(propTemplate.MethodTemplate)
		if err != nil {
			return "", "", nil, err
		}

		tmpl.Execute(methodCode, map[string]interface{}{
			"TypeKind":          tp.Kind(),
			"TypeName":          typeName,
			"MethodArguments":   fieldArguments,
			"MethodDescription": g.removeLineBreaks(g.returnString(fp.Description())),
			"MethodName":        name,
			"MethodReturnType":  fieldTypeName,
			"MethodReturn":      name,
			"Config":            conf,
			"TemplateConfig":    templateConfig,
		})

		imports = append(imports, g.getImports(fp.Type(), conf)...)
		imports = append(imports, propTemplate.Config.Imports...)
		imports = append(imports, propConf.Imports...)
		if val, ok := templateConfig["imports"]; ok {
			if arr, ok := val.([]string); ok {
				imports = append(imports, arr...)
			}
		}
	}

	return string(fieldCode.Bytes()), string(methodCode.Bytes()), imports, nil
}

func (g *CodeGen) getPointer(typeName string, fp *introspection.Field) string {
	if fp.Type().Kind() == "NON_NULL" {
		return typeName
	}
	return "*" + typeName
}

func (g *CodeGen) getImports(tp *introspection.Type, conf config.Config) []string {
	name := tp.Name()
	if name != nil {
		if val, ok := internalTypeConfig[*name]; ok {
			return []string{val.importPath}
		}
	}

	if tp.OfType() != nil {
		return g.getImports(tp.OfType(), conf)
	}

	return []string{}
}

func (g *CodeGen) getTypeName(tp *introspection.Type, conf config.Config, input bool) (typ string) {
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
		if len(typ) > 0 {
			if typ[len(typ)-1] != '*' {
				typ = typ + "*"
			}
		} else if tp.Kind() != "SCALAR" {
			typ = "*"
		}
		typ = typ + *name + "Resolver"
	} else {
		typ = typ + *name
	}

	if typ[0] != '*' && tp.Kind() == "INPUT_OBJECT" {
		typ = "*" + typ
	}

	return
}

func (g *CodeGen) removeDuplicates(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		val = strings.TrimSpace(val)
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func (g *CodeGen) isEntryPoint(a string) bool {
	return a == g.mutationName || a == g.queryName
}

func (g *CodeGen) removeLineBreaks(a string) string {
	return strings.Replace(a, "\n", " ", -1)
}

func (g *CodeGen) capitalise(str string) string {
	if strings.ToLower(str) == "id" {
		return "ID"
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}

func (g *CodeGen) unCapitalise(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}

func (g *CodeGen) subTemplate(str string, val interface{}) string {
	tmpl, err := template.New("sub_template").Funcs(g.templateFuncMap()).Parse(str)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	tmpl.Execute(buf, val)

	return string(buf.Bytes())
}

func (g *CodeGen) includesString(strings []string, str string) bool {
	for _, val := range strings {
		if val == str {
			return true
		}
	}

	return false
}

func (g *CodeGen) templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"capitalize":         g.capitalise,
		"uncapitalize":       g.unCapitalise,
		"is_entry":           g.isEntryPoint,
		"remove_line_breaks": g.removeLineBreaks,
		"sub_template":       g.subTemplate,
		"sprintf":            fmt.Sprintf,
		"includes_string":    g.includesString,
	}
}
