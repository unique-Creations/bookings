package render

import (
	"bytes"
	"fmt"
	"github.com/unique-Creations/bookings/config"
	"github.com/unique-Creations/bookings/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template cache.
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/templates.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from cache.
	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("Template not found in cache")
		return
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	//parseTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	//if err != nil {
	//	fmt.Println("error parsing templates:", err.Error())
	//	return
	//}
	//err = parseTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("error parsing templates:", err.Error())
	//	return
	//}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all the files in the templates directory
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// loop through the pages one-by-one
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// find all the layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

//var tc = make(map[string]*template.Template)
//
//func RenderTemplateTest(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	// Check if we already added template.
//
//	if _, found := tc[t]; !found {
//		// create template
//		log.Println("Creating template and adding to cache")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err.Error())
//		}
//	} else {
//		// we have the template in the cache
//	}
//
//	tmpl = tc[t]
//
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
//
//func createTemplateCache(t string) error {
//	tmplts := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//
//	tmpl, err := template.ParseFiles(tmplts...)
//	if err != nil {
//		return err
//	}
//
//	tc[t] = tmpl
//	return nil
//}
