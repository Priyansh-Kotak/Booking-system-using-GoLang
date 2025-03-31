package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Priyansh-Kotak/udemy-course-project/pkg/config"
)

// func RenderTempleteTest(w http.ResponseWriter, temp string) {
// 	parsedTemplete, err := template.ParseFiles("./templets/"+temp, "./templets/base.layout.html")
// 	if err != nil {
// 		log.Println("error while parsing templete")
// 		return
// 	}

// 	errs := parsedTemplete.Execute(w, nil)
// 	if errs != nil {
// 		log.Println("error while executing the parsed filr")
// 		return
// 	}
// }

// var tc = map[string]*template.Template
var app *config.AppConfig
// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplets(w http.ResponseWriter, t string) {
	//create a new templete cache
	// tc, err := CreateTemplateCache()
	tc := app.TemplateCache
	

	tmpl, errs := tc[t]
	if !errs {
		log.Fatal(errs)
	}

	buf := new(bytes.Buffer)

	err := tmpl.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

// Approach 1 for caching the records of parseFiles
// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templets/%s", t),
// 		"./templets/base.layout.html",
// 	}

// 	log.Println("printing templates ", templates)

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl
// 	log.Println("loging tmpl ", tmpl.ParseName)
// 	log.Println("pringitng ", tc)
// 	return nil
// }

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
