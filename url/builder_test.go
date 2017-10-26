package url

import (
	"testing"
)

const endpoint = "http://127.0.0.1"

func TestGetUrlByDate2000_01_01(t *testing.T){
	expect := endpoint + "/2000-01-01"
	actual := HistoryUrl(endpoint, 2000, 1, 1)
	if expect != actual {
		t.Error("expect " + expect + " actual " + actual)
	}
}
