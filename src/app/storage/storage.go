package storage

import (
	"fmt"
	"github.com/xiaosongluo/dashboard/src/app/database"
)

//DB interface for database
type Storage interface {
	Put(dashboardID string, data []byte) error
	Get(dashboardID string) ([]byte, error)
	Delete(dashboardID string) error
	Exists(dashboardID string) (bool, error)
}

//GetDatabase to get GetDatabase
func GetDatabase(db database.Database) (Storage, error) {
	switch db.Type {
	case "File":
		return NewFileStorage(db), nil
	}
	return nil, AdapterNotFoundError{db.Type}
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

//Error defines error
func (e DashboardNotFoundError) Error() string {
	return fmt.Sprintf("Dashboard with ID '%s' was not found.", e.DashboardID)
}
