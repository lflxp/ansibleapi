package api

import "testing"

func TestApi_Version(t *testing.T) {
	rs := Apis.Version()
	if len(rs) != 0 {
		t.Log(rs)
	} else {
		t.Error("rs is none")
	}
}

func TestApi_Setup(t *testing.T) {
	tmp := Apis.Setup("all")
	if len(tmp) > 0 {
		t.Log("setup result not none")
		if tmp["10.6.200.121"].Host != "" {
			t.Log(tmp["10.6.200.121"].Host)
			t.Log(tmp["10.6.200.121"].Status)
			t.Log(tmp["10.6.200.121"].Origin)
		} else {
			t.Error("10.6.200.121 is no available")
		}
	} else {
		t.Error("setup is none")
	}

}

func TestApi_ScriptA(t *testing.T) {
	tmp := Apis.ScriptA("10.6.200.121","/root/time.sh")
	if len(tmp) > 0 {
		t.Log("setup result not none")
		if tmp["10.6.200.121"].Host != "" {
			t.Log(tmp["10.6.200.121"].Host)
			t.Log(tmp["10.6.200.121"].Status)
			t.Log(tmp["10.6.200.121"].Origin)
		} else {
			t.Error("10.6.200.121 is no available")
		}
	} else {
		t.Error("setup is none")
	}
}

func TestApi_Copy(t *testing.T) {
	tmp := Apis.Copy("all","/root/time.sh","/tmp/time.sh","644","root")
	if len(tmp) > 0 {
		t.Log("setup result not none")
		if tmp["10.6.200.121"].Host != "" {
			t.Log(tmp["10.6.200.121"].Host)
			t.Log(tmp["10.6.200.121"].Status)
			t.Log(tmp["10.6.200.121"].Origin)
		} else {
			t.Error("10.6.200.121 is no available")
		}
	} else {
		t.Error("setup is none")
	}
}