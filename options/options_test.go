package options

import (
	"testing"
)

func TestOptions_NewOptions(t *testing.T) {
	rs := NewOptions("1","2","3","4")
	if rs.Type == "1" {
		t.Log("ok")
	} else {
		t.Error("not ok")
	}
}
