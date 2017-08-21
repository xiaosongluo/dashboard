package controllers

import (
	"github.com/flosch/pongo2"
	"net/http"
	"github.com/xiaosongluo/dashboard/src/app/view"
)

// DocumentHandller handle http request
func DocumentHandller(res http.ResponseWriter, req *http.Request) {

	view.RenderTemplate("document.html", pongo2.Context{}, res)
}
