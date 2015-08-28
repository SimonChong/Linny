package constants

import (
	"regexp"
)

const (
	AssetsDir   = "assets"
	MetricsDir  = "metrics"
	InsightsDir = "insights"
)

func AssetsRouteReg() *regexp.Regexp {
	return regexp.MustCompile(`^/` + AssetsDir + `/(?P<file>[^\.]+(?:\.html)?)$`)
}
