package main

import (
	"fmt"

	"github.com/simonchong/linny/common"

	"github.com/simonchong/linny/packer"
	"github.com/simonchong/linny/server"
)

func main() {

	flags := common.NewCmdFlags()

	configLinny, e1 := common.LoadConfigLinny()
	if e1 != nil {
		fmt.Println("ConfigLinny Error:", e1)
		return
	}

	if flags.Pack {
		packer.Pack(configLinny.ContentRoot)
	}
	if flags.Unpack != "" {
		packer.Unpack(configLinny, flags.Unpack)
	}

	if flags.Serve {

		configAd, e2 := common.LoadConfigAd(&configLinny)
		if e2 != nil {
			fmt.Println("ConfigAd Error:", e2)
			return
		}
		server.Start(configLinny, configAd)

	}
}
