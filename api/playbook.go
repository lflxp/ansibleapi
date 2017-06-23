package api

type Playbook struct {
	Api
}

var Playbooks Playbook

func init() {
	Playbooks = Playbook{}
	Playbooks.PlayBookInit()
}

func (this *Playbook) PlayBook(path,args string) string {
	this.Now = ""
	return this.SetOptions(path,args).ExecutePlayBookByRemote().GetResult()
}

func (this *Playbook) By() *Playbook {
	return this
}