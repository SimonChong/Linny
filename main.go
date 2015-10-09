package main

import (
	"flag"
	"fmt"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/creator"

	"github.com/simonchong/linny/insights"
	"github.com/simonchong/linny/packer"
	"github.com/simonchong/linny/server"
)

func main() {

	flags := common.NewCmdFlags()

	if flags.Init {
		creator.Create(flags.InitDir)
	}

	if flags.ConfigNeeded() {
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

			data := new(insights.Data)
			data.Init()
			defer data.Close()

			if e2 != nil {
				fmt.Println("ConfigAd Error:", e2)
				return
			}
			server.Start(&configLinny, &configAd, data)
		}

	}
	if flags.None() {
		flag.PrintDefaults()
	}

}
