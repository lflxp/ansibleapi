package api

type Playbook struct {
	Api
}

var Playbooks Playbook

func init() {
	Playbooks = Playbook{}
	Playbooks.PlayBookInit()
}

func (this *Playbook) PlayBook(host,path,args string) string {
	this.Now = ""
	return this.SetHosts(host).SetOptions(path,args).ExecutePlayBook().GetResult()
}