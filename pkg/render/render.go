package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AccurateApplication/go-web-app/pkg/config"
)

var funcMap = template.FuncMap{}

var conf *config.AppConf

// GenerateTemplates sets the config for the render package
func GenerateTemplates(c *config.AppConf) {
	conf = c
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tmplCache := conf.TemplateCache

	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatalf("Could not find template in map with value: %v\n", tmplCache[tmpl])
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

// Generate a cache as map with string and a pointer to template to render site with
func CreateTmplCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	layoutPath := "./templates/*.layout.tmpl"
	templatePath := "./templates/*.page.tmpl"

	pages, err := filepath.Glob(templatePath)
	if err != nil {
		return cache, err
	} else if len(pages) == 0 {
		log.Fatalf("Found no templates with glob: %s\n", templatePath)
	}

	for _, page := range pages {
		pageName := filepath.Base(page)

		// Get template set.
		ts, err := template.New(pageName).Funcs(funcMap).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob(layoutPath)
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob(layoutPath)
			if err != nil {
				return cache, err
			}
			cache[pageName] = ts
		}

	}
	return cache, nil
}
