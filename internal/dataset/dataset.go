package dataset

import (
	"encoding/json"
	"io"
	"os"
)

type City struct {
	Zone    int    `json:"zone"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Dataset struct {
	Cities [1440]*City `json:"cities"`
}

func Write(d Dataset, o io.Writer) error {
	enc := json.NewEncoder(o)
	enc.SetIndent("", "  ")
	if err := enc.Encode(d); err != nil {
		return err
	}
	return nil
}

func Read(f string) (Dataset, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return Dataset{}, err
	}
	var d Dataset
	if err := json.Unmarshal(b, &d); err != nil {
		return Dataset{}, err
	}

	return d, nil
}
