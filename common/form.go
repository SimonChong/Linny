package common

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var isNum = regexp.MustCompile(`\d+`)

func FormTime(field string, r *http.Request) (time.Time, error) {

	timeGen := r.FormValue("g")
	if !isNum.MatchString(timeGen) {
		return *new(time.Time), errors.New("Time value is not a number")
	}

	timeGenUnix, errT := strconv.ParseInt(timeGen, 10, 64)
	if errT != nil {
		return *new(time.Time), errors.New("Time value failed to parse")
	}

	now := time.Now()
	if timeGenUnix > now.Unix() {
		timeGenUnix = now.Unix()
	}
	if timeGenUnix < 0 {
		timeGenUnix = 0
	}

	return time.Unix(timeGenUnix, 0), nil
}
