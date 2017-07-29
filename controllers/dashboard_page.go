package controllers

import (
	"encoding/json"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	"github.com/xiaosongluo/dashboard/models"
	"net/http"
)

// GetDashboardHandller handle http request
func GetDashboardHandller(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	dash, err := models.LoadDashboard(params["dashid"], models.Database)
	if err != nil {
		dash = &models.Dashboard{APIKey: models.GenerateAPIKey(), Metrics: models.DashboardMetrics{}}
	}

	// Filter out expired metrics
	metrics := models.DashboardMetrics{}
	metrics = append(metrics, dash.Metrics...)

	renderTemplate("dashboard.html", pongo2.Context{
		"dashid":  params["dashid"],
		"metrics": metrics,
		"apikey":  dash.APIKey,
		"baseurl": models.Cfg.BaseURL,
	}, res)
}

// DeleteDashboardHandller handle http request
func DeleteDashboardHandller(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	dash, err := models.LoadDashboard(params["dashid"], models.Database)
	if err != nil {
		http.Error(res, "This dashboard does not exist.", http.StatusInternalServerError)
		return
	}

	if dash.APIKey != req.Header.Get("Authorization") {
		http.Error(res, "APIKey did not match.", http.StatusUnauthorized)
		return
	}

	models.Database.Delete(params["dashid"])
	http.Error(res, "OK", http.StatusOK)
}

// GetDashboardJsonHandler handle http request
func GetDashboardJsonHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	dash, err := models.LoadDashboard(params["dashid"], models.Database)
	if err != nil {
		dash = &models.Dashboard{APIKey: models.GenerateAPIKey(), Metrics: models.DashboardMetrics{}}
	}

	response := struct {
		APIKey  string `json:"api_key,omitempty"`
		Metrics []struct {
			ID          string  `json:"id"`
			Title       string  `json:"title"`
			Description string  `json:"description"`
			Label       string  `json:"label"`
			Status      string  `json:"status"`
			Value       float64 `json:"value"`
		} `json:"metrics"`
	}{}

	// Filter out expired metrics
	for _, m := range dash.Metrics {
		response.Metrics = append(response.Metrics, struct {
			ID          string  `json:"id"`
			Title       string  `json:"title"`
			Description string  `json:"description"`
			Label       string  `json:"label"`
			Status      string  `json:"status"`
			Value       float64 `json:"value"`
		}{
			ID:          m.MetricID,
			Title:       m.Title,
			Description: m.Description,
			Label:       m.Label,
			Status:      m.Status,
			Value:       m.Value,
		})
	}

	if len(response.Metrics) == 0 {
		response.APIKey = dash.APIKey
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}
