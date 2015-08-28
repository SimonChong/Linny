package constants

import (
	"regexp"
)

const (
	AssetsRoute   = "assets"
	MetricsRoute  = "metrics"
	InsightsRoute = "insights"
)

func AssetsRouteReg() *regexp.Regexp {
	return regexp.MustCompile(`^/` + AssetsRoute + `/(?P<file>[^\.]+(?:\.html)?)$`)
}
