package server

import (
	"github.com/zenazn/goji"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/constants"
	"github.com/simonchong/linny/insights"
	"github.com/simonchong/linny/server/controllers"
	"github.com/simonchong/linny/server/wrappers"
)

func Start(configLinny *common.ConfigLinny, configAd *common.ConfigAd, data *insights.Data) {

	appContext := wrappers.AppContext{
		Data:      data,
		ConfLinny: configLinny,
		ConfAd:    configAd,
	}

	goji.Get(constants.AssetsRouteReg(), wrappers.AppSessionHandler{&appContext, controllers.AssetHTML})
	goji.Get("/"+constants.AssetsDir+"/*", wrappers.AppSessionHandler{&appContext, controllers.AssetFiles})

	goji.Get("/"+constants.MeasureDir+"/click", wrappers.AppSessionHandler{&appContext, controllers.MeasureClick})

	goji.Get("/"+constants.ViewsDir+"/v.gif", wrappers.AppSessionHandler{&appContext, controllers.ViewCounter})

	goji.Serve()

}
