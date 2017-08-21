package controllers

import (
	"github.com/flosch/pongo2"
	"github.com/xiaosongluo/dashboard/src/app/view"
	"net/http"
)

// DocumentHandller handle http request
func DocumentHandller(res http.ResponseWriter, req *http.Request) {

	view.RenderTemplate("document.html", pongo2.Context{}, res)
}
