package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gobuffalo/plush/v4"
)

func currentPathWithout(values map[string]interface{}, ctx plush.HelperContext) string {
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
}
