package cmd

import (
	"fmt"
	"github.com/lflxp/ansibleapi/utils"
	//. "github.com/lflxp/ansibleapi/utils"
	"errors"
	"strings"
)

const (
	AnsibleCmd = "ssh root@10.6.200.131 ansible"
	PlayBookCmd = "ssh root@10.6.200.131 ansible-playbook"
	P = "-P"
	private_key = "--private-key"
	T = "-T"
	Scp_extra_args = "--scp-extra-args"
	B = "-B"
	D = "-D"
	L_l = "-l"
	List_hosts = "--list-hosts"
	Ssh_extra_args = "--ssh-extra-args"
	New_vault_password_file = "--new-vault-password-file"
	Ask_vault_pass = "--ask-vault-pass"
	L_f = "-f"
	L_h = "-h"
	L_m = "-m"
	Ask_sudo_pass = "--ask-sudo-pass"
	L_i = "-i"
	M = "-M"
	Version = "--version"
	U = "-U"
	L_o = "-o"
	Syntax_check = "--syntax-check"
	L_u = "-u"
	Become_user = "--become-user"
	Ask_su_pass = "--ask-su-pass"
	V = "-V"
	L_k = "-k"
	Sftp_extra_args = "--sftp-extra-args"
	R = "-R"
	S = "-S"
	K = "-K"
	L_a = "-a"
	C = "-C"
	Output = "--output"
	Ssh_common_args = "--ssh-common-args"
	L_b = "-b"
	Become_method = "--become-method"
	L_e = "-e"
	L_t = "-t"
	L_c = "-c"
	L_s = "-s"
)

//执行结构体
type Cmd struct {
	Command		string  //执行命令
	Host 		string 	//目标主机
	Now 		string //最终的命令
	Result 		string //结果
	IsSetHosts 	bool //是否设置了目标主机
}

//结果集结构体
type Result struct {
	Origin 		string //原始结果
	Host 		string //目标主机
	Status 		string 	//状态
}

//初始化 ansible命令执行结果
//初始化是否设置主机为否
func (this *Cmd) Init() {
	this.Command = AnsibleCmd
	this.IsSetHosts = false
}

//初始化 ansible命令执行结果
//初始化是否设置主机为否
func (this *Cmd) PlayBookInit() {
	this.Command = PlayBookCmd
	this.IsSetHosts = false
}

//设置选项
func (this *Cmd) SetOptions(key,value string) *Cmd {
	this.Now += fmt.Sprintf("%s %s ",key,value)
	return this
}

//设置COPY选项
func (this *Cmd) SetCopyOptions(key,src,dest,mode,owner string) *Cmd {
	this.Now += fmt.Sprintf("%s 'src=%s dest=%s mode=%s owner=%s'",key,src,dest,mode,owner)
	return this
}

//设置目标主机
func (this *Cmd) SetHosts(host string) *Cmd {
	this.Host = host
	this.IsSetHosts = true
	return this
}

//获取执行命令
func (this *Cmd) GetCmdByRemote() string {
	return fmt.Sprintf("%s %s \"%s\"",this.Command,this.Host,this.Now)
}

//获取执行命令
func (this *Cmd) GetCmdByLocal() string {
	return fmt.Sprintf("%s %s %s",this.Command,this.Host,this.Now)
}

//获取playbook执行命令
func (this *Cmd) GetPlayBookCmdByRemote() string {
	return fmt.Sprintf("%s '%s'",this.Command,this.Now)
}

//获取playbook执行命令
func (this *Cmd) GetPlayBookCmdByLocal() string {
	return fmt.Sprintf("%s %s",this.Command,this.Now)
}

//获取playbook执行结果
func (this *Cmd) GetResult() string {
	return this.Result
}

//执行命令
func (this *Cmd) ExecuteByRemote() *Cmd {
	if this.IsSetHosts == false {
		this.Result = "未设置目标主机"
		return this
	}
	this.Result = utils.ExecCommand(this.GetCmdByRemote())
	return this
}

//执行命令
func (this *Cmd) ExecuteByLocal() *Cmd {
	if this.IsSetHosts == false {
		this.Result = "未设置目标主机"
		return this
	}
	this.Result = utils.ExecCommand(this.GetCmdByLocal())
	return this
}


//执行playbook命令
func (this *Cmd) ExecutePlayBookByRemote() *Cmd {
	this.Result = utils.ExecCommand(this.GetPlayBookCmdByRemote())
	return this
}

//执行playbook命令
func (this *Cmd) ExecutePlayBookByLocal() *Cmd {
	this.Result = utils.ExecCommand(this.GetPlayBookCmdByLocal())
	return this
}

//解析结果 返回json数据
func (this *Cmd) ParseResult() (error,map[string]*Result) {
	result := map[string]*Result{}
	if this.IsSetHosts == false {
		return errors.New("未设置目标主机"),nil
	}

	origin_data := strings.Split(this.Result,"=>")
	for num,x := range origin_data {
		tmprs := &Result{}
		//println(num,x)
		if num + 2 <= len(origin_data) {
			tmp := strings.Split(x,"\n")
			tmp_status := strings.Split(tmp[len(tmp)-1]," | ")
			//println(tmp[len(tmp)-1],tmp_status)
			tmprs.Host = strings.TrimSpace(tmp_status[0])
			tmprs.Status = strings.TrimSpace(tmp_status[1])
			//按\n划分 去掉最后一个
			ttt := strings.Split(origin_data[num+1],"\n")
			tmprs.Origin = strings.Join(ttt[:len(ttt)-1],"\n")
			result[tmprs.Host] = tmprs
		}
	}

	return nil,result
}