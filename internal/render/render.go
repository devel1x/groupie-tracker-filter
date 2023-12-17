package render

import (
	"bytes"
	"fmt"
	"groupie/internal/config"
	"groupie/internal/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	funcstions = template.FuncMap{}
	app        *config.AppConfig
	err        error
)

// NewTemplate sets the config for render
func NewTemplate(a *config.AppConfig) {
	app = a
}

// RenderTemplate render template in html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Get the template cache from the app config
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
		fmt.Print("USED CASH")
	} else {
		tc, err = CreateTemplateCache() // error needs to be handled
		if err != nil {
			log.Fatal("coult not create cache")
		}
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("coult not get template from cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td) // error needs to be handled

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./web/templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(funcstions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		mathces, err := filepath.Glob("./web/templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(mathces) > 0 {
			ts, err = ts.ParseGlob("./web/templates/*layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, err
}
