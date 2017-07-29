package controllers

import (
	"net/http"
	"github.com/flosch/pongo2"
)

// HomeHandler handle http request
func HomeHandler(res http.ResponseWriter, req *http.Request) {

	renderTemplate("home.html", pongo2.Context{}, res)
}
