package main

import (
	"github.com/lflxp/ansibleapi/api"
)

func main() {
	//println(api.Apis.Version())
	tmp := api.Apis.Setup("all")
	println(tmp.Host)
	println(tmp.Origin)
}
