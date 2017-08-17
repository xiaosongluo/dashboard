package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xiaosongluo/dashboard/app/controllers"
	"github.com/xiaosongluo/dashboard/app/database"
	"github.com/xiaosongluo/dashboard/app/models"
	"github.com/xiaosongluo/dashboard/app/server"
	"github.com/xiaosongluo/dashboard/app/storage"
	"github.com/xiaosongluo/dashboard/app/utils/config"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var err error

	// Load the configuration file
	config.Load("config"+string(os.PathSeparator)+"config.json", cfg)

	_, err = controllers.PreloadTemplates()
	if err != nil {
		fmt.Printf("An error occurred while loading the templates handler: %s", err)
	}

	models.Storage, err = storage.GetDatabase(cfg.Database)
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

	addr := cfg.Server.Hostname + ":" + strconv.Itoa(cfg.Server.HTTPPort)
	http.ListenAndServe(addr, nil)
}

var cfg = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database database.Database `json:"Database"`
	Server   server.Server     `json:"Server"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
