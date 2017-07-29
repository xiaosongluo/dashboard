package db

import (
	"github.com/xiaosongluo/dashboard/conf"
	"fmt"
)

type DB interface {
	Put(dashboardID string, data []byte) error
	Get(dashboardID string) ([]byte, error)
	Delete(dashboardID string) error
	Exists(dashboardID string) (bool, error)
}

func GetDatabase(cfg *conf.Config) (DB, error) {
	switch cfg.Storage {
	case "file":
		return NewFileDatabase(cfg), nil
	}

	return nil, AdapterNotFoundError{cfg.Storage}
}

// AdapterNotFoundError is a named error for more simple determination which
// type of error is thrown
type AdapterNotFoundError struct {
	Name string
}

func (e AdapterNotFoundError) Error() string {
	return fmt.Sprintf("Storage '%s' not found.", e.Name)
}

// DashboardNotFoundError signalizes the requested dashboard could not be found
type DashboardNotFoundError struct {
	DashboardID string
}

func (e DashboardNotFoundError) Error() string {
	return fmt.Sprintf("Dashboard with ID '%s' was not found.", e.DashboardID)
}
