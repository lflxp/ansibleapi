package utils

import (
	"testing"
	"errors"
)

func TestCheck2(t *testing.T) {
	defer func(){
		if err := recover(); err != nil {
			t.Log(err)
		} else {
			t.Fatal("err is not nil and not err panic")
		}
	}()
	x := errors.New("test")
	Check(x)
}

func TestExecCommand(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
		}
	}()
	rs := ExecCommand("hostname")
	if len(rs) != 0 {
		t.Logf("成功 %s",rs)
	} else {
		t.Errorf("失败 返回结果为空")
	}
}