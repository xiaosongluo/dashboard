package config

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_Load_Success(t *testing.T) {
	Load("", cfg)
	t.Error("GenerateAPIKey failed")
}

var cfg = &configuration{}

type configuration struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

type mockedFilePath struct {
	reportErr     bool
	reportAbsPath string
}

func (m mockedFilePath) Abs(path string) (string, error) {
	if m.reportErr {
		return "", os.ErrNotExist
	}
	return m.reportAbsPath, nil
}

type mockedFile struct {
	// Embed this so we only need to add methods used by testable functions
	os.File
	fd   uintptr
	name string
}

type mockedIO struct {
	reportErr     bool
	reportErrType error
	reportFile    *os.File
}

func (m mockedIO) Open(name string) (*os.File, error) {
	if m.reportErr {
		return nil, m.reportErrType
	}
	return m.reportFile, nil
}
