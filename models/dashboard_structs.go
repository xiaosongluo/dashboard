package models

import (
	"encoding/json"
	"errors"
	"github.com/xiaosongluo/dashboard/storage"
	"github.com/xiaosongluo/dashboard/utils"
	"log"
	"strconv"
)

//Dashboard struct
type Dashboard struct {
	DashboardID string           `json:"-"`
	APIKey      string           `json:"api_key"`
	Metrics     DashboardMetrics `json:"metrics"`
	Storage     storage.Storage
}

//DashboardMetrics
type DashboardMetrics []*DashboardMetric

//DashboardMetric
type DashboardMetric struct {
	MetricID       string                 `json:"id"`
	Title          string                 `json:"title"`
	Description    string                 `json:"description"`
	LastLabel      string                 `json:"last_label"`
	LastStatus     string                 `json:"last_status"`
	LastValue      float64                `json:"last_value,omitifempty"`
	HistoricalData DashboardMetricHistory `json:"history,omitifempty"`
}

type DashboardMetricHistory []dashboardMetricStatus

type dashboardMetricStatus struct {
	Label  string  `json:"label"`
	Status string  `json:"status"`
	Value  float64 `json:"value"`
}

//LabelHistory
func (dm *DashboardMetric) LabelHistory() []string {
	s := []string{}

	labelStart := len(dm.HistoricalData) - 60
	if labelStart < 0 {
		labelStart = 0
	}

	for _, v := range dm.HistoricalData[labelStart:] {
		s = append(s, v.Label)
	}

	return s
}

//DataHistory
func (dm *DashboardMetric) DataHistory() []string {
	s := []string{}

	dataStart := len(dm.HistoricalData) - 60
	if dataStart < 0 {
		dataStart = 0
	}

	for _, v := range dm.HistoricalData[dataStart:] {
		s = append(s, strconv.FormatFloat(v.Value, 'g', 4, 64))
	}

	return s
}

//LoadDashboard
func LoadDashboard(dashid string, store storage.Storage) (*Dashboard, error) {
	data, err := store.Get(dashid)
	if err != nil {
		return &Dashboard{}, errors.New("Dashboard not found")
	}

	tmp := &Dashboard{
		DashboardID: dashid,
		Storage:     store,
	}
	_ = json.Unmarshal(data, tmp)

	return tmp, nil
}

//NewDashboardMetric
func NewDashboardMetric() *DashboardMetric {
	return &DashboardMetric{
		LastStatus:     "Unknown",
		HistoricalData: DashboardMetricHistory{},
	}
}

//IsValid
func (dm *DashboardMetric) IsValid() (bool, string) {
	if !utils.StringInSlice(dm.LastStatus, []string{"OK", "Warning", "Critical", "Unknowm"}) {
		return false, "Status not allowed"
	}

	if len(dm.Title) > 512 || len(dm.Description) > 1024 {
		return false, "Title or Description too long"
	}

	return true, ""
}

//Update
func (dm *DashboardMetric) Update(m *DashboardMetric) {
	dm.Title = m.Title
	dm.Description = m.Description
	dm.LastStatus = m.LastStatus
	dm.LastValue = m.LastValue
	dm.HistoricalData = append(dm.HistoricalData, dashboardMetricStatus{
		Label:  "\"" + m.LastLabel + "\"",
		Status: m.LastStatus,
		Value:  m.LastValue,
	})
}

//Save
func (d *Dashboard) Save() {
	data, err := json.Marshal(d)
	if err != nil {
		log.Printf("Error while marshalling dashboard: %s", err)
		return
	}
	err = d.Storage.Put(d.DashboardID, data)
	if err != nil {
		log.Printf("Error while storing dashboard: %s", err)
	}
}
