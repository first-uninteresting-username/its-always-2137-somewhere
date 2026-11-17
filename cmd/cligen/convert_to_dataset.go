package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

func countryName(code string) (string, error) {
	r, err := language.ParseRegion(code)
	if err != nil {
		return "", err
	}
	if !r.IsCountry() {
		return "", fmt.Errorf("%q is not a country", code)
	}
	return display.English.Regions().Name(r), nil
}
