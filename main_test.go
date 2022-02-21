package main

import (
	"os"
	"testing"
	"time"

	"github.com/kyokomi/emoji/v2"
)

func TestGetWorld(t *testing.T) {
	tstStr := GetWorldTxt()
	refStr := emoji.Sprint(":worldmap:")
	if tstStr != refStr {
		t.Errorf("%s it's not a worldmap pic ", tstStr)
	}
}

func TestWillTimeout(t *testing.T) {
	time.Sleep(1 * time.Second)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
