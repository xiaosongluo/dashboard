package main

import (
	"encoding/json"
	"fmt"
	"github.com/xiaosongluo/dashboard/src/app/database"
	"github.com/xiaosongluo/dashboard/src/app/models"
	"github.com/xiaosongluo/dashboard/src/app/route"
	"github.com/xiaosongluo/dashboard/src/app/server"
	"github.com/xiaosongluo/dashboard/src/app/storage"
	"github.com/xiaosongluo/dashboard/src/app/utils/config"
	"github.com/xiaosongluo/dashboard/src/app/view"
	"os"
)

func main() {

	var err error

	// Load the configuration file
	config.Load("config"+string(os.PathSeparator)+"config.json", cfg)

	_, err = view.PreloadTemplates()
	if err != nil {
		fmt.Printf("An error occurred while loading the templates handler: %s", err)
	}

	models.Storage, err = storage.GetDatabase(cfg.Database)
	if err != nil {
		fmt.Printf("An error occurred while loading the storage handler: %s", err)
	}

	server.Run(route.LoadHTTP(), route.LoadHTTPS(), cfg.Server)
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
