package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/xiaosongluo/dashboard/controllers"
	"fmt"
	"github.com/xiaosongluo/dashboard/db"
	"github.com/xiaosongluo/dashboard/conf"
	"github.com/xiaosongluo/dashboard/models"
)

var (
)

func main() {

	var err error

	_, err = controllers.PreloadTemplates()
	if err != nil {
		fmt.Printf("An error occurred while loading the storage handler: %s", err)
	}

	models.Cfg = conf.Load()
	models.Database, err = db.GetDatabase(models.Cfg)
	if err != nil {
		fmt.Printf("An error occurred while loading the storage handler: %s", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")

	router.HandleFunc("/doc", controllers.DocumentHandller).Methods("GET")

	router.HandleFunc("/{dashid}.json", controllers.GetDashboardJsonHandler).Methods("GET")

	router.HandleFunc("/{dashid}", controllers.GetDashboardHandller).Methods("GET")
	router.HandleFunc("/{dashid}", controllers.DeleteDashboardHandller).Methods("DELETE")

	router.HandleFunc("/{dashid}/{metricid}", controllers.PutMetricHandller).Methods("PUT")
	router.HandleFunc("/{dashid}/{metricid}", controllers.DeleteMetricHandller).Methods("DELETE")

	router.HandleFunc("/{dashid}.json", controllers.GetDashboardJsonHandler).Methods("GET")

	/*router.HandleFunc("/{dashid}/{metricid}.json", handleDisplayMetricJSON).Methods("GET")*/

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./statics/")))

	http.Handle("/", router)
	http.ListenAndServe(":3001", nil)
}
