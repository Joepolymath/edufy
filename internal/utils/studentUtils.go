package utils

import "errors"

func ValidateQueryStatus(status string) error {
	// List of allowed statuses
	allowedStatuses := []string{"ENROLLED", "FAILED", "COMPLETED"}

	// Check if the provided status is in the allowed statuses
	for _, s := range allowedStatuses {
		if s == status {
			return nil // Status is valid
		}
	}

	// If status is not found in the list, return an error
	return errors.New("Invalid status")
}
