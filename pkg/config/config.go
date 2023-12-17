package config

import "text/template"

// Hold the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
