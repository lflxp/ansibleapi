package api

import (
	. "github.com/lflxp/ansibleapi/cmd"
	"github.com/lflxp/ansibleapi/model"
	. "github.com/lflxp/ansibleapi/utils"
)

type Api struct {
	Cmd
}

var Apis Api

func init() {
	Apis = Api{}
	Apis.Init()
}

func (this *Api) Version() string {
	this.IsSetHosts = true
	this.Now = ""
	this.SetOptions(Version,"").Execute()
	return this.Result
}

func (this *Api) Setup(host string) *Result {
	this.Now = ""
	err,data := this.SetHosts(host).SetOptions(L_m,model.SETUP).Execute().ParseResult()
	Check(err)
	return data
}