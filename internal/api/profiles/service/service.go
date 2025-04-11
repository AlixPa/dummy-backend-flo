package service

import (
	"log"
	"os"
	"path"

	"github.com/gocarina/gocsv"
)

type Profile struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

type Config struct {
	DbPath string
}

type Service struct {
	dbPath string
}

func New(cfg Config) *Service {
	return &Service{dbPath: cfg.DbPath}
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

func (s *Service) ListProfiles() []*Profile {
	ls, err := loadProfiles(path.Join(s.dbPath, "profiles.csv"))
	log.Print(err)
	return ls
}
