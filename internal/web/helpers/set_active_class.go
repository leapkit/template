package helpers

import (
	"net/http"
	"regexp"

	"github.com/gobuffalo/plush/v4"
)

func setActiveClass(pattern []interface{}, active, inactive string, ctx plush.HelperContext) string {
	req, ok := ctx.Value("request").(*http.Request)
	if !ok {
		return inactive
	}

	pathName := req.URL.Path
	for _, p := range pattern {
		re, err := regexp.Compile(p.(string))
		if err != nil {
			return inactive
		}

		if re.MatchString(pathName) {
			return active
		}
	}

	return inactive
}
