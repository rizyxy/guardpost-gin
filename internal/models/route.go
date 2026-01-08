package models

type Route struct {
	Path       string `yaml:"path" json:"path"`
	Method     string `yaml:"method" json:"method"`
	Downstream string `yaml:"downstream" json:"downstream"`
}
