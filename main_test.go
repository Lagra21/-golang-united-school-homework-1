package main

import (
	"os"
	"testing"
	"time"

	"github.com/kyokomi/emoji/v2"
)

func TestGetWorld(t *testing.T) {
	testRes := "ok"
	tstStr := GetWorldTxt()
	refStr := emoji.Sprint(":map:")
	if tstStr != refStr {
		testRes = "fail"
		t.Errorf("%s it's not a worldmap pic ", tstStr)
	}

	t.Logf(testRes)
}

func TestTimeout(t *testing.T) {
	time.Sleep(1 * time.Second)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
