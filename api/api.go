package api

import (
	. "github.com/lflxp/ansibleapi/cmd"
	"github.com/lflxp/ansibleapi/model"
	. "github.com/lflxp/ansibleapi/utils"
)

type Api struct {
	Cmd
}

//初始化 ansible命令执行结果
//初始化是否设置主机为否
func (this *Api) Init() {
	this.Command = AnsibleCmd
	this.IsSetHosts = false
}

func (this *Api) Version() string {
	this.IsSetHosts = true
	this.Now = ""
	this.SetOptions(Version,"").ExecuteByRemote()
	return this.Result
}

func (this *Api) Setup(host string) map[string]*Result {
	this.Now = ""
	err,data := this.SetHosts(host).SetOptions(L_m,model.SETUP).ExecuteByRemote().ParseResult()
	Check(err)
	return data
}

func (this *Api) ScriptA(host,path string) map[string]*Result {
	this.Now = ""
	err,data := this.SetHosts(host).SetOptions(L_m,model.SCRIPT).SetOptions(L_a,path).ExecuteByRemote().ParseResult()
	Check(err)
	return data
}

func (this *Api) Copy(host,src,dest,mode,owner string) map[string]*Result {
	this.Now = ""
	err,data := this.SetHosts(host).SetOptions(L_m,model.COPY).SetCopyOptions(L_a,src,dest,mode,owner).ExecuteByRemote().ParseResult()
	Check(err)
	return data
}