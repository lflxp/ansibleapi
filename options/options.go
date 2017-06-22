package options

import "fmt"

/**
ansible命令配置管理
 */

var optionList map[string]*Options

type Options struct {
	Cmd 		string
	Short 		string
	Common 		string
	Type 		string
	UserInput 	string
}

func NewOptions(types,cmd,short,common string) *Options {
	opt := &Options{Type:types,Cmd:cmd,Short:short,Common:common}
	return opt
}

func SetOptions(key,value string) *Options {
	result := optionList[key]
	result.UserInput = value
	return result
}

func GetOption(key string) *Options {
	return optionList[key]
}

func init() {
	optionList = map[string]*Options{}
	//Options
	optionList["-a"] = NewOptions("Options","--args","-a","MODULE_ARGS, --args=MODULE_ARGS module arguments")
	optionList["--ask-vault-pass"] = NewOptions("Options","--ask-vault-pass ","","ask for vault password")
	optionList["-B"] = NewOptions("Options","--background","-B","SECONDS, --background=SECONDS")
	optionList["-C"] = NewOptions("Options","--check","-C","don't make any changes; instead, try to predict some of the changes that may occur")
	optionList["-D"] = NewOptions("Options","--diff","-D","when changing (small) files and templates, show the differences in those files; works great with --check")
	optionList["-e"] = NewOptions("Options","--extra-vars","-e","EXTRA_VARS, --extra-vars=EXTRA_VARS")
	optionList["-f"] = NewOptions("Options","--forks","-f","FORKS, --forks=FORKS specify number of parallel processes to use (default=5)")
	optionList["-h"] = NewOptions("Options","--help","-h","show this help message and exit")
	optionList["-i"] = NewOptions("Options","--inventory-file","-i","INVENTORY, --inventory-file=INVENTORY (default=/etc/ansible/hosts) or comma separated host list")
	optionList["-l"] = NewOptions("Options","--limit","-l","show this help message and exit")
	optionList["--list-hosts"] = NewOptions("Options","--list-hosts","","outputs a list of matching hosts; does not execute anything else")
	optionList["-m"] = NewOptions("Options","--module-name","-m","MODULE_NAME, --module-name=MODULE_NAME \n module name to execute (default=command)")
	optionList["-M"] = NewOptions("Options","--module-path","-M","MODULE_PATH, --module-path=MODULE_PATH")
	optionList["--new-vault-password-file"] = NewOptions("Options","--new-vault-password-file","","new vault password file for rekey")
	optionList["-o"] = NewOptions("Options","--one-line","","condense output")
	optionList["--output"] = NewOptions("Options","--output","","output file name for encrypt or decrypt; use - for stdout")
	optionList["-P"] = NewOptions("Options","--poll","-P","POLL_INTERVAL, --poll=POLL_INTERVAL")
	optionList["--syntax-check"] = NewOptions("Options","--syntax-check","","perform a syntax check on the playbook, but do not execute it")
	optionList["-t"] = NewOptions("Options","--tree","-t","TREE, --tree=TREE  log output to this directory")
	optionList["-V"] = NewOptions("Options","--verbose","-V","verbose mode (-vvv for more, -vvvv to enable connection debugging)")
	optionList["--version"] = NewOptions("Options","--version","","show program's version number and exit")

	//Connection Options
	optionList["-k"] = NewOptions("Connection Options","--ask-pass","-k","ask for connection password")
	optionList["--private-key"] = NewOptions("Connection Options","--private-key","--key-file=","ask for connection password")
	optionList["-u"] = NewOptions("Connection Options","--user","-u","REMOTE_USER, --user=REMOTE_USER connect as this user (default=None)")
	optionList["-c"] = NewOptions("Connection Options","--connection","-c","CONNECTION, --connection=CONNECTION")
	optionList["-T"] = NewOptions("Connection Options","--timeout","-T","TIMEOUT, --timeout=TIMEOUT override the connection timeout in seconds \n default=10)")
	optionList["--ssh-common-args"] = NewOptions("Connection Options","--ssh-common-args","","specify common arguments to pass to sftp/scp/ssh (e.g. ProxyCommand)")
	optionList["--sftp-extra-args"] = NewOptions("Connection Options","--sftp-extra-args","","specify extra arguments to pass to sftp only (e.g. -f, -l)")
	optionList["--scp-extra-args"] = NewOptions("Connection Options","--scp-extra-args","","specify extra arguments to pass to scp only (e.g. -l)")
	optionList["--ssh-extra-args"] = NewOptions("Connection Options","--ssh-extra-args","","specify extra arguments to pass to ssh only (e.g. -R)")

	//Privilege Escalation Options:
	optionList["-s"] = NewOptions("Privilege Escalation Options","--sudo","-s","run operations with sudo (nopasswd) (deprecated, use become)")
	optionList["-S"] = NewOptions("Privilege Escalation Options","--sudo","--su","run operations with su (deprecated, use become)")
	optionList["-U"] = NewOptions("Privilege Escalation Options","--sudo-user","-U","SUDO_USER, --sudo-user=SUDO_USER desired sudo user (default=root) (deprecated, use become)")
	optionList["-R"] = NewOptions("Privilege Escalation Options","--su-user","-R","SU_USER, --su-user=SU_USER run operations with su as this user (default=root) (deprecated, use become)")
	optionList["-b"] = NewOptions("Privilege Escalation Options","--become","-b","run operations with become (does not imply password prompting)")
	optionList["--become-method"] = NewOptions("Privilege Escalation Options","--become-method","","privilege escalation method to use (default=sudo), valid choices: [ sudo | su | pbrun | pfexec | doas | dzdo | ksu | runas ]")
	optionList["--become-user"] = NewOptions("Privilege Escalation Options","--become-user","","run operations as this user (default=root)")
	optionList["--ask-sudo-pass"] = NewOptions("Privilege Escalation Options","--ask-sudo-pass","","ask for sudo password (deprecated, use become)")
	optionList["--ask-su-pass"] = NewOptions("Privilege Escalation Options","--ask-su-pass","","ask for su password (deprecated, use become)")
	optionList["-K"] = NewOptions("Privilege Escalation Options","--ask-become-pass","-K","ask for privilege escalation password")
	for k,_ := range optionList {
		fmt.Println(fmt.Sprintf("%s = \"%s\"",k,k))
	}
}
