package service

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/AlixPa/dummy-backend-flo/internal/common"
	"github.com/gocarina/gocsv"
)

type Profile struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

type Service struct {
	cfg *common.ProfileConfig
}

func New(cfg *common.ProfileConfig) *Service {
	return &Service{cfg}
}

func loadProfiles(path string) ([]*Profile, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	profiles := []*Profile{}
	if err := gocsv.UnmarshalFile(f, &profiles); err != nil {
		if errors.Is(err, gocsv.ErrEmptyCSVFile) {
			return profiles, nil
		}
		return nil, err
	}
	return profiles, nil
}

func (s *Service) ListProfiles() ([]*Profile, error) {
	return loadProfiles(s.cfg.DbCsvPath)
}

func profileExists(path string, name string) (bool, error) {
	profiles, err := loadProfiles(path)
	if err != nil {
		return true, nil
	}

	for _, p := range profiles {
		if p != nil && strings.EqualFold(p.Name, name) {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) CreateProfile(name string, age int) error {
	o, err := profileExists(s.cfg.DbCsvPath, name)
	if err != nil {
		return err
	}
	if o {
		return fmt.Errorf("Profile with name %s is already in database : %w", name, common.ErrDuplicateProfileName)
	}

	f, err := os.OpenFile(s.cfg.DbCsvPath, os.O_WRONLY|os.O_APPEND, 0644)
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
