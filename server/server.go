package server

import (
	"github.com/zenazn/goji"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/constants"
	"github.com/simonchong/linny/insights"
	"github.com/simonchong/linny/server/controllers"
)

func Start(configLinny *common.ConfigLinny, configAd *common.ConfigAd, data *insights.Data) {
	controllerFact := controllers.Factory{
		ConfLinny: configLinny,
		ConfAd:    configAd,
		Data:      data,
	}
	goji.Get(constants.AssetsRouteReg(), controllerFact.AssetHTML())
	goji.Get("/"+constants.AssetsDir+"/*", controllerFact.AssetFiles())

	goji.Get("/"+constants.MeasureDir+"/click", controllerFact.MeasureClick())

	goji.Serve()

}
