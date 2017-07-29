package db

import (
	"fmt"
	"github.com/xiaosongluo/dashboard/conf"
	"io/ioutil"
	"os"
	"path"
)

//FileDatabase
type FileDatabase struct {
	cfg *conf.Config
}

//NewFileDatabase new a FileDatabase
func NewFileDatabase(cfg *conf.Config) *FileDatabase {
	// Create directory if not exists
	if _, err := os.Stat(cfg.FileDatabase.Directory); os.IsNotExist(err) {
		if err := os.MkdirAll(cfg.FileDatabase.Directory, 0700); err != nil {
			fmt.Printf("Could not create storage directory '%s'", cfg.FileDatabase.Directory)
			os.Exit(1)
		}
	}

	return &FileDatabase{
		cfg: cfg,
	}
}

// Put writes the given data to FS
func (f *FileDatabase) Put(dashboardID string, data []byte) error {
	err := ioutil.WriteFile(f.getFilePath(dashboardID), data, 0600)

	return err
}

// Get loads the data for the given dashboard from FS
func (f *FileDatabase) Get(dashboardID string) ([]byte, error) {
	data, err := ioutil.ReadFile(f.getFilePath(dashboardID))
	if err != nil {
		return nil, DashboardNotFoundError{dashboardID}
	}

	return data, nil
}

// Delete deletes the given dashboard from FS
func (f *FileDatabase) Delete(dashboardID string) error {
	if exists, err := f.Exists(dashboardID); err != nil || !exists {
		if err != nil {
			return err
		}
		return DashboardNotFoundError{dashboardID}
	}

	return os.Remove(f.getFilePath(dashboardID))
}

// Exists checks for the existence of the given dashboard
func (f *FileDatabase) Exists(dashboardID string) (bool, error) {
	if _, err := os.Stat(f.getFilePath(dashboardID)); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (f *FileDatabase) getFilePath(dashboardID string) string {
	return path.Join(f.cfg.FileDatabase.Directory, dashboardID+".json")
}
