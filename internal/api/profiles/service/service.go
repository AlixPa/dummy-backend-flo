package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AlixPa/dummy-backend-flo/internal/api/profiles/model"
	"github.com/gocarina/gocsv"
	"golang.org/x/exp/slices"
)

type Service struct {
	cfg ServiceConfig
}

type ServiceConfig interface {
	GetProfilesTablePath() string
}

func New(cfg ServiceConfig) *Service {
	return &Service{cfg}
}

func loadProfiles(path string) ([]*model.Profile, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	profiles := []*model.Profile{}
	if err := gocsv.UnmarshalFile(f, &profiles); err != nil {
		if errors.Is(err, gocsv.ErrEmptyCSVFile) {
			return profiles, nil
		}
		return nil, err
	}
	return profiles, nil
}

func saveProfiles(path string, profiles []*model.Profile) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return gocsv.MarshalFile(&profiles, f)
}

func (s *Service) ListProfiles() ([]*model.Profile, error) {
	return loadProfiles(s.cfg.GetProfilesTablePath())
}

func (s *Service) GetProfile(id string) (*model.Profile, error) {
	profiles, err := loadProfiles(s.cfg.GetProfilesTablePath())
	if err != nil {
		return nil, err
	}

	for _, p := range profiles {
		if p != nil && p.ID == id {
			return p, nil
		}
	}
	return nil, ErrProfileNotFound
}

func profileExists(path string, name string) (bool, error) {
	profiles, err := loadProfiles(path)
	if err != nil {
		return true, nil
	}

	for _, p := range profiles {
		if p != nil && strings.EqualFold(*(p.Name), name) {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) CreateProfile(profile model.ProfileFields) error {
	if err := profile.FullFields(); err != nil {
		return err
	}

	if err := profile.Validate(); err != nil {
		return err
	}

	o, err := profileExists(s.cfg.GetProfilesTablePath(), *(profile.Name))
	if err != nil {
		return err
	}
	if o {
		return fmt.Errorf("profile with name %s is already in database: %w", *(profile.Name), ErrDuplicateProfileName)
	}

	profiles, err := loadProfiles(s.cfg.GetProfilesTablePath())
	if err != nil {
		return err
	}

	// Generate a new ID
	newID := strconv.Itoa(len(profiles) + 1)
	newProfile := &model.Profile{
		ID:            newID,
		ProfileFields: profile,
	}

	profiles = append(profiles, newProfile)
	return saveProfiles(s.cfg.GetProfilesTablePath(), profiles)
}

func (s *Service) UpdateProfile(id string, profile model.ProfileFields) error {
	if err := profile.Validate(); err != nil {
		return err
	}

	profiles, err := loadProfiles(s.cfg.GetProfilesTablePath())
	if err != nil {
		return err
	}

	found := false
	for _, p := range profiles {
		if p != nil && p.ID == id {
			if profile.Name != nil {
				p.Name = profile.Name
			}
			if profile.Age != nil {
				p.Age = profile.Age
			}
			found = true
			break
		}
	}

	if !found {
		return ErrProfileNotFound
	}

	return saveProfiles(s.cfg.GetProfilesTablePath(), profiles)
}

func (s *Service) DeleteProfile(id string) error {
	profiles, err := loadProfiles(s.cfg.GetProfilesTablePath())
	if err != nil {
		return err
	}

	found := false
	for i, p := range profiles {
		if p != nil && p.ID == id {
			profiles = slices.Delete(profiles, i, i+1)
			found = true
			break
		}
	}

	if !found {
		return ErrProfileNotFound
	}

	return saveProfiles(s.cfg.GetProfilesTablePath(), profiles)
}

var (
	ErrDuplicateProfileName = errors.New("duplicate profile name")
	ErrProfileNotFound      = errors.New("profile not found")
)
