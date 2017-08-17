package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xiaosongluo/dashboard/app/controllers"
	"github.com/xiaosongluo/dashboard/app/models"
	"github.com/xiaosongluo/dashboard/app/storage"
	"github.com/xiaosongluo/dashboard/app/utils/config"
	"net/http"
	"os"
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

	http.Handle("/", handlers.LoggingHandler(os.Stdout, router))
	http.ListenAndServe(models.Config.Listen, nil)
}
