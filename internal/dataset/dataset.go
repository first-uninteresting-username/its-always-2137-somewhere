package dataset

import (
	"encoding/json"
	"io"
)

type City struct {
	Zone    int    `json:"zone"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Dataset struct {
	Cities [1440]*City
}

func Write(d Dataset, o io.Writer) error {
	enc := json.NewEncoder(o)
	enc.SetIndent("", "  ")
	if err := enc.Encode(d); err != nil {
		return err
	}
	return nil
}
