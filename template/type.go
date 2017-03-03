package template

import (
	"errors"
	"path"

	"github.com/hashicorp/hcl"
)

//go:generate go-bindata -pkg template -ignore=\.go  -o assets.go ./...

type TypeTemplateConfig struct {
	Type    string
	Imports []string
}

type TypeTemplate struct {
	Config       TypeTemplateConfig
	TypeTemplate string
}

func parseTypeConfig(str string) (config TypeTemplateConfig, err error) {
	err = hcl.Decode(&config, str)
	return
}

func loadStoredTypeTemplate(template string) (*TypeTemplate, bool, error) {
	cfgStr, err := Asset(template + "/config.hcl")
	if err != nil {
		return nil, false, nil
	}

	cfg, err := parseTypeConfig(string(cfgStr))
	if err != nil {
		return nil, false, err
	}

	templateString, err := Asset(path.Join(template, cfg.Type))
	if err != nil {
		return nil, false, err
	}

	return &TypeTemplate{
		Config:       cfg,
		TypeTemplate: string(templateString),
	}, true, nil
}

func GetTypeTemplate(templateName string) (template *TypeTemplate, err error) {
	template, ok, err := loadStoredTypeTemplate("type/" + templateName)
	// if ok false read file from the directory
	if !ok {
		return nil, errors.New("Fail")
	}
	return
}
