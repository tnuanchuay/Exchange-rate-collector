package dtime

import (
	"testing"
	"fmt"
)

func TestGetTimeShouldReturn2000_01_01(t *testing.T){
	expect := "2000-01-01"
	actual := DateString(2000, 1, 1)

	if expect != actual {
		t.Error(fmt.Sprintf("expect %s actaul %s", expect, actual))
	}
}

func TestGetTimeShouldReturn2000_13_13(t *testing.T){
	expect := "2000-13-13"
	actual := DateString(2000, 13, 13)

	if expect != actual{
		t.Error("expect %s actual %s", expect, actual)
	}
}
