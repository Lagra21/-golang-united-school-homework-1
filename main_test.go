package main

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestShowWorld(t *testing.T) {
	tstStr := GetWorldTxt()
	refStr := emoji.Sprint(":map:")
	if tstStr != refStr {
		t.Errorf("%s it's not a correct pic ", tstStr)
	}
}
