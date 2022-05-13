package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConf holds our application config
type AppConf struct {
	TemplateCache map[string]*template.Template
	SecureSite    bool
	Session       *scs.SessionManager
}
