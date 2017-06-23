package main

import (
	"github.com/lflxp/ansibleapi/api"
	//"strings"
	//. "github.com/lflxp/ansibleapi/cmd"
)

func main() {
	//println(api.Apis.Version())
	//tmp := api.Apis.Setup("all")
	//tmp := api.Apis.ScriptA("10.6.200.121","/root/time.sh")
	//tmp := api.Apis.Copy("all","/root/time.sh","/tmp/time.sh","644","root")
	//println(tmp["10.6.200.121"].Host)
	//println(tmp["10.6.200.121"].Status)
	//println(tmp["10.6.200.121"].Origin)

	//tmp := api.Playbooks.PlayBook("all","/root/init.yml","-v")
	//println(tmp)


	tmp := api.Playbooks.By().SetOptions("/root/init.yml","-v").SetOptions("--extra-vars","\"hosts=10.6.200.121 user=root\"").ExecutePlayBookByRemote().GetResult()
	println(tmp)
}
