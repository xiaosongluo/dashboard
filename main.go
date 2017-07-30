package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xiaosongluo/dashboard/config"
	"github.com/xiaosongluo/dashboard/controllers"
	"github.com/xiaosongluo/dashboard/models"
	"github.com/xiaosongluo/dashboard/storage"
	"github.com/xiaosongluo/dashboard/utils"
	"net/http"
)

func main() {

	var err error

	_, err = controllers.PreloadTemplates()
	if err != nil {
		fmt.Printf("An error occurred while loading the storage handler: %s", err)
	}

	models.Config = config.Load()
	models.Storage, err = storage.GetDatabase(models.Config)
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

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./statics/")))

	http.Handle("/", utils.LogHTTPRequest(router))
	http.ListenAndServe(models.Config.Listen, nil)
}
