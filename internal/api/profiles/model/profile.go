package model

import (
	"errors"
	"fmt"
	"regexp"
)

// ProfileFields contains the field definitions and validation rules
type ProfileFields struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

// Profile represents a user profile in the database
type Profile struct {
	ID string `json:"id" csv:"id"`
	ProfileFields
}

// UpdateProfileRequest represents the request body for updating a profile
type UpdateProfileRequest struct {
	ProfileFields
}

// ProfileResponse represents the response for profile operations
type ProfileResponse struct {
	Data *Profile `json:"data,omitempty"`
}

// ProfilesResponse represents the response for listing profiles
type ProfilesResponse struct {
	Data []*Profile `json:"data"`
}

// ValidationError represents a collection of validation errors
type ValidationError struct {
	Errors []error
}

func (ve *ValidationError) Error() string {
	if len(ve.Errors) == 0 {
		return "no validation errors"
	}
	return fmt.Sprintf("validation errors: %v", ve.Errors)
}

func (ve *ValidationError) Add(err error) {
	if err != nil {
		ve.Errors = append(ve.Errors, err)
	}
}

func (ve *ValidationError) HasErrors() bool {
	return len(ve.Errors) > 0
}

var (
	ErrNameMissing = errors.New("missing Name for profile creation")
	ErrAgeMissing  = errors.New("missing Age for profile creation")
	ErrNameFormat  = errors.New("name field should be of size 2 to 50 characters")
	ErrAgeValue    = errors.New("age should be in between 0 and 150 years old")
	ErrNameValue   = errors.New("name should contain only letters")
)

func (p *ProfileFields) Validate() error {
	validationErr := &ValidationError{}

	if p.Name != nil {
		if len(*p.Name) < 2 || len(*p.Name) > 50 {
			validationErr.Add(ErrNameFormat)
		}
		validName := regexp.MustCompile(`^[a-zA-Z]+$`)
		if !validName.MatchString(*p.Name) {
			validationErr.Add(ErrNameValue)
		}
	}

	if p.Age != nil && (*p.Age < 0 || *p.Age > 150) {
		validationErr.Add(ErrAgeValue)
	}

	if validationErr.HasErrors() {
		return validationErr
	}

	return nil
}

func (p *ProfileFields) FullFields() error {
	validationErr := &ValidationError{}

	if p.Name == nil {
		validationErr.Add(ErrNameMissing)
	}
	if p.Age == nil {
		validationErr.Add(ErrAgeMissing)
	}

	if validationErr.HasErrors() {
		return validationErr
	}

	return nil
}
