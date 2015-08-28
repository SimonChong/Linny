package main

import (
	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/constants"
	"github.com/simonchong/linny/controllers"

	"github.com/zenazn/goji"
)

func main() {

	config := common.NewConfig()
	controllerFact := controllers.Factory{Conf: config}

	goji.Get(constants.AssetsRouteReg(), controllerFact.AssetHTML())
	goji.Get("/"+constants.AssetsDir+"/*", controllerFact.AssetFiles())

	goji.Serve()
}
