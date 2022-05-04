package config

import "html/template"

// AppConf holds our application config
type AppConf struct {
	TemplateCache map[string]*template.Template
}
