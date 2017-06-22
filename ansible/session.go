package ansible

import (
	. "github.com/widuu/gojson"
	. "github.com/lflxp/ansibleapi/utils"
)

//具体操作实现
//继承全局函数实现再调用
type Session struct {
	Ansible 	*Ansible //继承父结构
	IsPlayBook	bool
	Statement	*Statement //执行options模块
	OriginRs 	string	//原始结果
	BeforeRs	map[string]*Js //解析前的结果
	AfterRs		map[string]*Js //解析后的结果
}

func (this *Session) Clone() *Session {
	var sess = *this
	return &sess
}

func (this *Session) Init() {
	this.Statement.Init()
	this.Statement.Ansible = this.Ansible
	this.IsPlayBook = false
	this.OriginRs = ""

	this.BeforeRs = make(map[string]*Js)
	this.AfterRs = make(map[string]*Js)
}

func (this *Session) Version() *Session {
	this.Statement.SetOption("--version")
	println(this.Statement.AnsibleCmd)
	return this
}

func (this *Session) Execute() string {
	this.OriginRs = ExecCommand(this.Statement.AnsibleCmd)
	return this.OriginRs
}