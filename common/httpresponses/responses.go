package httpresponses

import "fmt"

type ValidationError struct {
	ErrorCode    int    `json:"errorCode"`    // What is displayed to the end-user on the frontend
	ErrorMessage string `json:"errorMessage"` // Helps frontend associate errors with UI elements
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %d", e.ErrorMessage, e.ErrorCode)
}

// Used in 20X responses.
type Response[T any] struct {
	Data             T                  `json:"data"`
	ValidationErrors []*ValidationError `json:"validationErrors"`
}

// Used in non-20X responses.
type ErrorResponse struct {
	ErrorCode string `json:"ErrorCode"`
}
