package constants

import (
	"regexp"
)

const (
	AssetsDir  = "assets"
	MeasureDir = "measure"
	ViewsDir   = "views"
)

func AssetsRouteReg() *regexp.Regexp {
	return regexp.MustCompile(`^/` + AssetsDir + `/(?P<file>[^\.]+(?:\.html)?)$`)
}
