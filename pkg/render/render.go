package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Priyansh-Kotak/udemy-course-project/pkg/config"
	"github.com/Priyansh-Kotak/udemy-course-project/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplets(w http.ResponseWriter, t string, td *models.TemplateData) {

	var tc map[string]*template.Template
	// render ever time the page refreshes
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}

	tmpl, errs := tc[t]
	if !errs {
		log.Fatal(errs)
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := tmpl.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the pages from the directory which has filename start from ./templets/*.page.html
	pages, err := filepath.Glob("./templets/*.page.html")
	if err != nil {
		return myCache, err
	}

	// now loop through each file and parse them
	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("printing name ", name, "page = ", page)

		ts, err := template.New(name).ParseFiles(page)
		log.Println("Printing ts ", ts)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templets/*.layout.html")
		log.Println("printing matches ", matches)
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templets/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
