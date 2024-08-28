package api

import "desafio-backend/pkg/errors"

type ErrorLink map[string]string

type Error struct {
	errors.SimpleError
	Links  ErrorLink `json:"links"`
	Status string    `json:"status"`
	Code   string    `json:"code"`
}

func (error *Error) WithLinkAbout(link string) Error {
	error.Links = ErrorLink{
		"about": link,
	}

	return *error
}

type Errors struct {
	Errors []Error `json:"errors"`
}

func (apiErrors *Errors) ToAPIErrors(status string, link string, ers ...errors.Error) *Errors {
	for _, err := range ers {
		apiError := Error{
			Links: ErrorLink{
				"about": link,
			},
			Status:      status,
			SimpleError: err.Error(),
		}

		apiErrors.Errors = append(apiErrors.Errors, apiError)
	}

	return apiErrors
}
