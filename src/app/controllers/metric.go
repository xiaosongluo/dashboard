package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xiaosongluo/dashboard/src/app/models"
	"io/ioutil"
	"net/http"
)

// PutMetricHandller handle http request
func PutMetricHandller(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	metricUpdate := models.NewDashboardMetric()
	err = json.Unmarshal(body, metricUpdate)
	if err != nil {
		http.Error(res, "Unable to unmarshal json", http.StatusInternalServerError)
		return
	}

	dash, err := models.LoadDashboard(params["dashid"], models.Storage)
	if err != nil {
		if len(req.Header.Get("Authorization")) < 10 {
			http.Error(res, "APIKey is too insecure", http.StatusUnauthorized)
			return
		}
		dash = &models.Dashboard{
			APIKey:      req.Header.Get("Authorization"),
			Metrics:     models.DashboardMetrics{},
			DashboardID: params["dashid"],
			Storage:     models.Storage,
		}
	}

	if dash.APIKey != req.Header.Get("Authorization") {
		http.Error(res, "APIKey did not match.", http.StatusUnauthorized)
		return
	}

	valid, reason := metricUpdate.IsValid()
	if !valid {
		http.Error(res, fmt.Sprintf("Invalid data: %s", reason), http.StatusInternalServerError)
		return
	}

	updated := false
	for _, m := range dash.Metrics {
		if m.MetricID == params["metricid"] {
			m.Update(metricUpdate)
			updated = true
			break
		}
	}

	if !updated {
		tmp := models.NewDashboardMetric()
		tmp.MetricID = params["metricid"]
		tmp.Update(metricUpdate)
		dash.Metrics = append(dash.Metrics, tmp)
	}

	dash.Save()

	http.Error(res, "OK", http.StatusOK)
}

// DeleteMetricHandller handle http request
func DeleteMetricHandller(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	dash, err := models.LoadDashboard(params["dashid"], models.Storage)
	if err != nil {
		dash = &models.Dashboard{APIKey: req.Header.Get("Authorization"), Metrics: models.DashboardMetrics{}, DashboardID: params["dashid"]}
	}

	if dash.APIKey != req.Header.Get("Authorization") {
		http.Error(res, "APIKey did not match.", http.StatusUnauthorized)
		return
	}

	tmp := models.DashboardMetrics{}
	for _, m := range dash.Metrics {
		if m.MetricID != params["metricid"] {
			tmp = append(tmp, m)
		}
	}
	dash.Metrics = tmp
	dash.Save()

	http.Error(res, "OK", http.StatusOK)
}
