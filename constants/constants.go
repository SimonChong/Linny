package constants

import (
	"regexp"
)

const (
	AssetsDir  = "assets"
	MeasureDir = "m"
)

func AssetsRouteReg() *regexp.Regexp {
	return regexp.MustCompile(`^/` + AssetsDir + `/(?P<file>[^\.]+(?:\.html)?)$`)
}
