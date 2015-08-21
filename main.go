package main

import (
	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/controllers"

	"github.com/zenazn/goji"
)

func main() {

	config := common.NewConfig()
	controllerFact := controllers.Factory{Conf: config}

	goji.Get("/*", controllerFact.AdHtml())

	goji.Serve()
}
