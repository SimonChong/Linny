package constants

import (
	"regexp"
)

const (
	AssetsDir   = "assets"
	MeasureDir  = "measure"
	InsightsDir = "insights"
)

func AssetsRouteReg() *regexp.Regexp {
	return regexp.MustCompile(`^/` + AssetsDir + `/(?P<file>[^\.]+(?:\.html)?)$`)
}
