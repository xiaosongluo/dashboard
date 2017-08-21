package controllers

import (
	"github.com/flosch/pongo2"
	"net/http"
	"github.com/xiaosongluo/dashboard/src/app/view"
)

// HomeHandler handle http request
func HomeHandler(res http.ResponseWriter, req *http.Request) {

	view.RenderTemplate("home.html", pongo2.Context{}, res)
}
