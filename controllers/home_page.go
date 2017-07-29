package controllers

import (
	"github.com/flosch/pongo2"
	"net/http"
)

// HomeHandler handle http request
func HomeHandler(res http.ResponseWriter, req *http.Request) {

	renderTemplate("home.html", pongo2.Context{}, res)
}
