package service

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Profile struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

type Config struct {
	DbCsvPath string
}

type Service struct {
	dbCsvPath string
}

func New(cfg Config) *Service {
	return &Service{dbCsvPath: cfg.DbCsvPath}
}

func loadProfiles(path string) ([]*Profile, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var profiles []*Profile
	if err := gocsv.UnmarshalFile(f, &profiles); err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *Service) ListProfiles() ([]*Profile, error) {
	return loadProfiles(s.dbCsvPath)
}

func (s *Service) CreateProfile(name string, age int) error {
	f, err := os.OpenFile(s.dbCsvPath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%s,%d\n", name, age)
	if err != nil {
		return err
	}

	return nil
}
