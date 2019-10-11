package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// MatchURL check if url matches pathname
// /material/{\d+} => /material/11
func MatchURL(pathname, url string) bool {
	var segs = strings.Split(pathname, "/")
	var paths = strings.Split(url, "/")
	if len(segs) != len(paths) {
		return false
	}
	for k, v := range segs {
		if strings.HasPrefix(v, "{") {
			var exp = v[1 : len(v)-1]
			var reg, err = regexp.Compile(fmt.Sprintf("^%s$", exp))
			if err != nil {
				return false
			}
			if !reg.MatchString(paths[k]) {
				return false
			}
		} else {
			if v != paths[k] {
				return false
			}
		}
	}
	return true
}
