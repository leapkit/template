package helpers

import (
	"net/http"

	"github.com/gobuffalo/plush/v4"
)

func currentPathWithoutFields(fields []interface{}, ctx plush.HelperContext) string {
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
}
