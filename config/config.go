package config

type FieldConfig struct {
	Template map[string]map[string]interface{}
	Imports  []string
}

type TypeConfig struct {
	Template map[string]map[string]interface{}
	Field    map[string]FieldConfig
	Imports  []string
}

type Config struct {
	Package string
	Type    map[string]TypeConfig
}
