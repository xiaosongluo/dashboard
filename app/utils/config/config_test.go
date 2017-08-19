package config

import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

func Test_Load_Success(t *testing.T) {
	filepath := &mockedFilePath{}
	io := &mockedIO{}
	ioutil := mockedIoutil{}

	filepath.reportErr = false
	filepath.reportAbsPath = "D:\test.json"
	io.reportErr = false
	io.reportErrType = nil
	io.reportFile = nil
	ioutil.reportErr = false
	ioutil.reportErrType = nil
	ioutil.reportByte = []byte(`{"id":1,"name":"cat"}`)

	Load("test.json", cfg)
	if cfg.Id != 1 || cfg.Name != "cat" {
		t.Error("Load module information wrong!")
	}
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

type mockedIoutil struct {
	reportErr     bool
	reportErrType error
	reportByte    []byte
}

func (m mockedIoutil) ReadAll(r io.Reader) ([]byte, error) {
	if m.reportErr {
		return nil, m.reportErrType
	}
	return m.reportByte, nil
}
