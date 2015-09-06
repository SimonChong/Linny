package common

import (
	"flag"
)

type CmdFlags struct {
	Pack   bool
	Unpack string
	Serve  bool
}

func NewCmdFlags() CmdFlags {
	f := CmdFlags{}
	f.Collect()
	return f
}

func (a *CmdFlags) Collect() {
	serve := flag.Bool("serve", false, "Start the server. Serve the ad or campaign in the folder specified in configLinny.json")
	pack := flag.Bool("pack", false, "Pack current ad or campaign into an .adpack file")
	unpack := flag.String("unpack", "", "UnPack specified .adpack file and update configLinny.json point to it")
	flag.Parse()

	a.Serve = *serve
	a.Pack = *pack
	a.Unpack = *unpack
}
