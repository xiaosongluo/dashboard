package config

import (
	"encoding/json"
	"testing"
	"github.com/spf13/afero"
)

func Test_Load_Success(t *testing.T) {

	appFS := afero.NewOsFs()

	appFS.MkdirAll("fakedir/tmp", 0755)
	afero.WriteFile(appFS, "fakedir/tmp/test.json", []byte(`{"id":1,"name":"cat"}`), 0777)
	name := "fakedir/tmp/test.json"

	Load(name, cfg)
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