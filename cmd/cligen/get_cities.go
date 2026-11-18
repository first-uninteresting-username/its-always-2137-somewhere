package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type rawCity struct {
	Name        string
	CountryCode string
	Longitude   float64
	Population  int
	Zone        *int
}

type rawCities struct {
	Cities []rawCity
}

func getCities(path string) (rawCities, error) {
	f, err := os.Open(path)
	if err != nil {
		return rawCities{}, err
	}
	defer f.Close()

	var rc rawCities
	sc := bufio.NewScanner(f)
	for n := 1; sc.Scan(); n++ {
		fields := strings.Split(sc.Text(), "\t")
		name, countryCode := fields[1], fields[8]
		longitude, err := strconv.ParseFloat(fields[5], 64)
		if err != nil {
			return rawCities{}, err
		}
		population, err := strconv.Atoi(fields[14])
		if err != nil {
			return rawCities{}, err
		}

		c := rawCity{
			Name:        name,
			CountryCode: countryCode,
			Longitude:   longitude,
			Population:  population,
			Zone:        nil,
		}

		rc.Cities = append(rc.Cities, c)
	}

	if err := sc.Err(); err != nil {
		return rawCities{}, err
	}

	return rc, nil
}
