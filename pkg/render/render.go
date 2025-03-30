package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTempleteTest(w http.ResponseWriter, temp string) {
	parsedTemplete, err := template.ParseFiles("./templets/"+temp, "./templets/base.layout.html")
	if err != nil {
		log.Println("error while parsing templete")
		return
	}

	errs := parsedTemplete.Execute(w, nil)
	if errs != nil {
		log.Println("error while executing the parsed filr")
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplets(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, isMap := tc[t]

	if !isMap {
		// we will create a new templete
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we will take it from the cache
		log.Println("taking from the cache")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// Approach 1 for caching the records of parseFiles
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templets/%s", t),
		"./templets/base.layout.html",
	}

	log.Println("printing templates ", templates)

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = tmpl
	log.Println("loging tmpl ", tmpl.ParseName)
	log.Println("pringitng ", tc)
	return nil
}
