package controllers

import (
	"net/http"
	"github.com/flosch/pongo2"
)

func DocumentHandller(res http.ResponseWriter, req *http.Request) {

	renderTemplate("document.html", pongo2.Context{}, res)
}
