package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/first-uninteresting-username/its-always-2137-somewhere/internal/dataset"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var emptyZoneError = errors.New("Zone is empty")

const EMPTY_CITY string = "EMPTY_CITY"

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

func getBiggestInZone(zone int, d rawCities) (rawCity, error) {
	c := rawCity{
		Name:       EMPTY_CITY,
		Population: 0,
	}
	for _, city := range d.Cities {
		if *city.Zone == zone {
			if city.Population > c.Population {
				c = city
			}
		}
	}
	if c.Name == EMPTY_CITY {
		return rawCity{}, emptyZoneError
	}

	return c, nil
}

func addZones(cities rawCities) (rawCities, error) {
	var citi rawCities
	for _, c := range cities.Cities {
		var zone *int
		*zone = getZone(c.Longitude)
		c = rawCity{
			Name:        c.Name,
			CountryCode: c.CountryCode,
			Longitude:   c.Longitude,
			Population:  c.Population,
			Zone:        zone,
		}
		citi.Cities = append(citi.Cities, c)
	}
	return citi, nil
}

func convertToDataset(cities rawCities) (dataset.Dataset, error) {
	var d dataset.Dataset
	for zone := range 1339 {
		c, err := getBiggestInZone(zone, cities)
		if err != nil {
			return dataset.Dataset{}, err
		}
		country, err := countryName(c.CountryCode)
		if err != nil {
			return dataset.Dataset{}, err
		}
		city := dataset.City{
			Zone:    zone,
			Name:    c.Name,
			Country: country,
		}

		d.Cities[zone] = &city
	}

	return d, nil
}
