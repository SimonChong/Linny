package server

import (
	"regexp"

	"github.com/zenazn/goji"

	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/insights"
	"github.com/simonchong/linny/server/controllers"
	"github.com/simonchong/linny/server/paths"
	"github.com/simonchong/linny/server/wrappers"
)

func Start(configLinny *common.ConfigLinny, configAd *common.ConfigAd, data *insights.Data) {

	appContext := wrappers.AppContext{
		Data:      data,
		ConfLinny: configLinny,
		ConfAd:    configAd,
	}

	goji.Get("/"+paths.MeasureDir+"/k", wrappers.AppSessionHandler{AppContext: &appContext, Handler: controllers.ClickTracking})

	goji.Get("/"+paths.MeasureDir+"/v.gif", wrappers.AppSessionHandler{AppContext: &appContext, Handler: controllers.ViewCounter})

	goji.Get("/"+paths.MeasureDir+"/c.js", wrappers.AppContextHandler{AppContext: &appContext, Handler: controllers.ConversionsJS})

	goji.Get("/"+paths.MeasureDir+"/c.gif", wrappers.AppContextHandler{AppContext: &appContext, Handler: controllers.Conversions})

	goji.Get(assetsRouteReg(), wrappers.AppSessionHandler{AppContext: &appContext, Handler: controllers.AssetHTML})
	goji.Get("/*", wrappers.AppSessionHandler{AppContext: &appContext, Handler: controllers.AssetFiles})

	goji.Serve()

}

func assetsRouteReg() *regexp.Regexp {
	return regexp.MustCompile(`^/(?P<file>[^\.]+(?:\.html)?)$`)
}
