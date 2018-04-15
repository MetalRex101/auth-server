package validators

import (
	"net/url"
)

type UrlValidator struct {
	Base *Validator
}

func (v UrlValidator) compareUrls (url1 string, url2 string) (bool, error) {
	u1, err := url.Parse(url1)

	if err != nil { return false, err }

	u2, err := url.Parse(url2)

	if err != nil { return false, err }

	if u1.Scheme != u2.Scheme { return false, nil }

	if u1.Host != u2.Host { return false, nil }

	if u1.Port() != u2.Port() { return false, nil }

	return true, nil
}