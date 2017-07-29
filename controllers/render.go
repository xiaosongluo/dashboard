package controllers

import (
	"net/http"
	"fmt"
	"github.com/flosch/pongo2"
	"io/ioutil"
)

var (
	templates = make(map[string]*pongo2.Template)
)

// PreloadTemplates finds the templates files
func PreloadTemplates() (n int, err error) {
	templateNames, err := ioutil.ReadDir("templates")
	if err != nil {
		panic("Templates directory not available!")
	}
	for _, tplname := range templateNames {
		templates[tplname.Name()] = pongo2.Must(pongo2.FromFile(fmt.Sprintf("templates/%s", tplname.Name())))
	}
	return len(templates), err
}

func renderTemplate(templateName string, context pongo2.Context, res http.ResponseWriter) {
	if tpl, ok := templates[templateName]; ok {
		_ = tpl.ExecuteWriter(context, res)
	} else {
		res.WriteHeader(http.StatusInternalServerError)
		_, _ = res.Write([]byte(fmt.Sprintf("Template %s not found!", templateName)))
	}
}
