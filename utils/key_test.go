package utils

import (
	"testing"
)

func Test_GenerateAPIKey_Success(t *testing.T) {
	key := GenerateAPIKey()
	if len(key) == 0 {
		t.Error("GenerateAPIKey failed")
	}
}
