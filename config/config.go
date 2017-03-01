package config

type FieldConfig struct {
	Template map[string]map[string]interface{}
}

type TypeConfig struct {
	Template map[string]map[string]interface{}
	Field    map[string]FieldConfig
}

type Config struct {
	Package string
	Type    map[string]TypeConfig
}
