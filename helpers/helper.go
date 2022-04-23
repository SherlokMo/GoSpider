package helpers

import (
	"errors"
	"net/url"
)

func HnadleError(err error) {
	if err != nil {
		panic(err)
	}
}

func ProvisionURL(toProvision, base string) (string, error) {
	var err error = nil

	if len(toProvision) < 1 || toProvision[0] == '#' {
		err = errors.New("Empty Hyperlink")
	} else if toProvision[0] == '/' {
		toProvision = base + toProvision
	} else if toProvision[0:2] == "./" {
		toProvision = getBaseUrl(base) + toProvision[1:]
	}

	return toProvision, err
}

func getBaseUrl(t string) string {
	u, _ := url.Parse(t)

	return u.Scheme + "://" + u.Host
}
