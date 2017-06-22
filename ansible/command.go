package ansible

import (
	. "github.com/lflxp/ansibleapi/options"
)

const (
	AnsibleCmd = "ssh root@10.6.200.131 ansible "
)

type Statement struct {
	Ansible 	*Ansible
	Option 		*Options
	AnsibleCmd 	string
}

func (this *Statement) Init() {
	this.AnsibleCmd = AnsibleCmd
}

func (this *Statement) SetOption(key ...string) *Statement {
	if len(key) == 2 {
		this.Option = SetOptions(key[0],key[1])
		if this.Option.Short != "" {
			this.AnsibleCmd += this.Option.Short+" "+this.Option.UserInput
		} else {
			this.AnsibleCmd += this.Option.Cmd+" "+this.Option.UserInput
		}
	} else if len(key) == 1 {
		this.Option = SetOptions(key[0],"")
		if this.Option.Short != "" {
			this.AnsibleCmd += this.Option.Short
		} else {
			this.AnsibleCmd += this.Option.Cmd
		}
	}

	return this
}