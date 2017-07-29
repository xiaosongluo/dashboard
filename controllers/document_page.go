package controllers

import (
	"github.com/flosch/pongo2"
	"net/http"
)

// DocumentHandller handle http request
func DocumentHandller(res http.ResponseWriter, req *http.Request) {

	renderTemplate("document.html", pongo2.Context{}, res)
}
