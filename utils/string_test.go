package utils

import (
	"testing"
)

func Test_StringInSlice_Suceess(t *testing.T) {
	if !StringInSlice("Warning", []string{"OK", "Warning", "Critical", "Unknowm"}) {
		t.Error("Warning should in [OK, Warning, Critical, Unknowm]")
	}
}

func Test_StringInSlice_Fail(t *testing.T) {
	if StringInSlice("Nothing", []string{"OK", "Warning", "Critical", "Unknowm"}) {
		t.Error("Nothing should NOT in [OK, Warning, Critical, Unknowm]")
	}
}
