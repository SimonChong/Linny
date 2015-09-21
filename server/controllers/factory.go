package controllers

import (
	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/insights"
)

type Factory struct {
	Data      *insights.Data
	ConfLinny *common.ConfigLinny
	ConfAd    *common.ConfigAd
}
