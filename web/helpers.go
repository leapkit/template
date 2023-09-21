package web

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gobuffalo/plush/v4"
)

// Helpers used by the application to render the templates.
var helpers = map[string]any{
	// "form":    form.Form,
	// "formFor": tags.FormFor,

	"hasPrefix": func(str, prefix string) bool {
		return strings.HasPrefix(str, prefix)
	},

	"urlWithParam": func(path string, params map[string]any) string {
		u, err := url.Parse(path)
		if err != nil {
			return path
		}

		q := u.Query()
		for k, v := range params {
			q.Set(k, fmt.Sprintf("%v", v))
		}

		u.RawQuery = q.Encode()
		return u.String()
	},

	"urlHas": func(path, key, val string) bool {
		u, err := url.Parse(path)
		if err != nil {
			return false
		}

		q := u.Query()
		return q.Get(key) == val
	},

	"currentPathWithoutFields": func(fields []interface{}, ctx plush.HelperContext) string {
		req, ok := ctx.Value("request").(*http.Request)
		if !ok {
			return ""
		}

		path := req.URL
		q := path.Query()
		for _, field := range fields {
			q.Del(field.(string))
		}

		path.RawQuery = q.Encode()
		return path.String()
	},

	"currentPathWithout": func(values map[string]interface{}, ctx plush.HelperContext) string {
		req, ok := ctx.Value("request").(*http.Request)
		if !ok {
			return ""
		}

		URL, err := url.QueryUnescape(req.URL.String())
		if err != nil {
			return ""
		}

		for key, val := range values {
			URL = strings.ReplaceAll(URL, fmt.Sprintf("%s=%s", key, val), "")
		}

		URL = strings.ReplaceAll(URL, "&&", "&")

		return URL
	},

	"ifURLMatches": func(pattern string, matches, notMatch any, ctx plush.HelperContext) any {
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
	},

	"setActiveClass": func(pattern []interface{}, active, inactive string, ctx plush.HelperContext) string {
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
	},
}
