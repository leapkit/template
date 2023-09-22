package helpers

import (
	"net/http"
	"regexp"

	"github.com/gobuffalo/plush/v4"
)

func urlMatches(pattern string, matches, notMatch any, ctx plush.HelperContext) any {
	req, ok := ctx.Value("request").(*http.Request)
	if !ok {
		return notMatch
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return notMatch
	}

	if re.MatchString(req.URL.Path) {
		return matches
	}

	return notMatch
}
