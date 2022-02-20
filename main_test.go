package main

import (
	"os"
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestGetWorld(t *testing.T) {
	tstStr := GetWorldTxt()
	refStr := emoji.Sprint(":map:")
	if tstStr != refStr {
		t.Errorf("%s it's not a correct pic ", tstStr)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
