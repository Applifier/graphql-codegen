package template

import (
	"errors"

	"github.com/hashicorp/hcl"
)

//go:generate go-bindata -pkg template -ignore=\.go  -o assets.go ./...

type PropertyTemplateConfig struct {
	Field  string
	Method string
}

type PropertyTemplate struct {
	Config         PropertyTemplateConfig
	FieldTemplate  string
	MethodTemplate string
}

func parseConfig(str string) (config PropertyTemplateConfig, err error) {
	err = hcl.Decode(&config, str)
	return
}

func loadStoredTemplate(template string) (*PropertyTemplate, bool, error) {
	cfgStr, err := Asset(template + "/config.hcl")
	if err != nil {
		return nil, false, nil
	}

	cfg, err := parseConfig(string(cfgStr))
	if err != nil {
		return nil, false, err
	}

	fieldTemplateString, err := Asset(template + "/" + cfg.Field)
	if err != nil {
		return nil, false, err
	}

	methodTemplateString, err := Asset(template + "/" + cfg.Method)
	if err != nil {
		return nil, false, err
	}

	return &PropertyTemplate{
		Config:         cfg,
		FieldTemplate:  string(fieldTemplateString),
		MethodTemplate: string(methodTemplateString),
	}, true, nil
}

func GetPropertyTemplate(templateName string) (template *PropertyTemplate, err error) {
	template, ok, err := loadStoredTemplate("property/" + templateName)
	// if ok false read file from the directory
	if !ok {
		return nil, errors.New("Fail")
	}
	return
}
