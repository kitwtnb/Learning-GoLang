package gosample_test

import (
	"testing"
	"gosample"
)

func TestExapmle(t *testing.T) {
	s := gosample.Message
	if s == "" {
		t.Fail()
	}
}
