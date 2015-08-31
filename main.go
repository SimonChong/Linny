package main

import (
	"fmt"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/constants"
	"github.com/simonchong/linny/controllers"

	"github.com/zenazn/goji"
)

func main() {

	configLinny, e1 := common.LoadConfigLinny()
	if e1 != nil {
		fmt.Println("ConfigLinny Error:", e1)
		return
	}
	configAd, e2 := common.LoadConfigAd(&configLinny)
	if e2 != nil {
		fmt.Println("ConfigAd Error:", e2)
		return
	}

	controllerFact := controllers.Factory{
		ConfLinny: configLinny,
		ConfAd:    configAd,
	}

	goji.Get(constants.AssetsRouteReg(), controllerFact.AssetHTML())
	goji.Get("/"+constants.AssetsDir+"/*", controllerFact.AssetFiles())

	goji.Get("/"+constants.MetricsDir+"/click", controllerFact.MetricsClick())

	goji.Serve()
}
