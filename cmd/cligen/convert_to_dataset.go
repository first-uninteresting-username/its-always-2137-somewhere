package main

import (
	"fmt"
	"math"

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

func getZone(longitude float64) int {
	// I want zones to be counted from east to west, but longitudes are rising eastward
	// That is why I need to reverse them

	long := (-longitude + 180) * 4
	return int(math.Floor(long))
}
