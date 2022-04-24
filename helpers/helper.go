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
		toProvision = getBaseUrl(base) + toProvision
	} else if toProvision[0:2] == "./" {
		toProvision = getBaseUrl(base) + toProvision[1:]
	}

	err = isValidProtocol(toProvision)
	return toProvision, err
}

func getBaseUrl(t string) string {
	u, _ := url.Parse(t)

	return u.Scheme + "://" + u.Host
}

func isValidProtocol(t string) error {
	u, err := url.ParseRequestURI(t)
	if err != nil {
		return err
	}

	if u.Scheme == "http" || u.Scheme == "https" {
		return nil
	}
	return errors.New("Not valid web protocol")
}
