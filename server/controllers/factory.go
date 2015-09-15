package controllers

import (
	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/insights"
)

type Factory struct {
	ConfLinny *common.ConfigLinny
	ConfAd    *common.ConfigAd
	Data      *insights.Data
}
