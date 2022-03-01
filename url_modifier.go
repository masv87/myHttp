package main

import (
	urls "net/url"
)

type UrlModifier struct {
}

func NewUrlModifier() *UrlModifier {
	return &UrlModifier{}
}

func (uM *UrlModifier) Modify(address string) (string, error) {
	url, err := urls.Parse(address)
	if err != nil {
		return "", err
	}
	url.Scheme = "http"
	return url.String(), nil
}
