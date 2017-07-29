package controllers

import (
	"net/http"
	"github.com/flosch/pongo2"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {

	renderTemplate("home.html", pongo2.Context{}, res)
}
