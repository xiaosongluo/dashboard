package utils

import (
	"testing"
)

func Test_CheckStringInSlice_Success(t *testing.T) {
	if !CheckStringInSlice("Warning", []string{"OK", "Warning", "Critical", "Unknowm"}) {
		t.Error("Warning should in [OK, Warning, Critical, Unknowm]")
	}
}

func Test_CheckStringInSlice_Fail(t *testing.T) {
	if CheckStringInSlice("Nothing", []string{"OK", "Warning", "Critical", "Unknowm"}) {
		t.Error("Nothing should NOT in [OK, Warning, Critical, Unknowm]")
	}
}
