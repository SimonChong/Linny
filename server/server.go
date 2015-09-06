package server

import (
	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/constants"
	"github.com/simonchong/linny/server/controllers"

	"github.com/zenazn/goji"
)

func Start(configLinny common.ConfigLinny, configAd common.ConfigAd) {
	controllerFact := controllers.Factory{
		ConfLinny: configLinny,
		ConfAd:    configAd,
	}
	goji.Get(constants.AssetsRouteReg(), controllerFact.AssetHTML())
	goji.Get("/"+constants.AssetsDir+"/*", controllerFact.AssetFiles())

	goji.Get("/"+constants.MetricsDir+"/click", controllerFact.MetricsClick())

	goji.Serve()
}
