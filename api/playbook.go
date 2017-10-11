package api

import . "github.com/lflxp/ansibleapi/cmd"

type Playbook struct {
	Api
}

//初始化 ansible命令执行结果
//初始化是否设置主机为否
func (this *Playbook) PlayBookInit() {
	this.Command = PlayBookCmd
	this.IsSetHosts = false
}
func (this *Playbook) PlayBook(path,args string) string {
	this.Now = ""
	this.SetOptions(path,args).ExecutePlayBookByRemote()
	return this.GetResult()
}

func (this *Playbook) By() *Playbook {
	return this
}