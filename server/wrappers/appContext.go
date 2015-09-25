package wrappers

import (
	"github.com/simonchong/linny/common"
	"github.com/simonchong/linny/insights"
)

type AppContext struct {
	Data      *insights.Data
	ConfLinny *common.ConfigLinny
	ConfAd    *common.ConfigAd
}
