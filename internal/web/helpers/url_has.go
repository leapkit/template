package helpers

import "net/url"

func urlHas(path, key, val string) bool {
	u, err := url.Parse(path)
	if err != nil {
		return false
	}

	q := u.Query()
	return q.Get(key) == val
}
