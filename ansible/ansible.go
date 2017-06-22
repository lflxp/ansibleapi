package ansible

import (
	. "github.com/widuu/gojson"
)

//全局结构 包含所有方法
type Ansible struct {
	OriginRs 	string	//原始结果
	BeforeRs	map[string]*Js //解析前的结果
	AfterRs		map[string]*Js //解析后的结果
}

func (this *Ansible) NewSession() *Session {
	session := &Session{Ansible:this}
	session.Init()
	//println(session.OriginRs)
	return session
}

func (this *Ansible) Version() string {
	session := this.NewSession()
	session.Version()
	return session.Execute()
}