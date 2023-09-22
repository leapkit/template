package helpers

import (
	"fmt"
	"net/url"
)

func urlWithParams(path string, params map[string]any) string {
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
}
