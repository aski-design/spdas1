package email

import (
	"errors"
	"strings"
)

const atSymbol = "@"

var (
	ErrMissingFrom      = errors.New("email: sender address is required")
	ErrInvalidFrom      = errors.New("email: sender address is invalid")
	ErrMissingRecipient = errors.New("email: at least one recipient is required")
	ErrInvalidRecipient = errors.New("email: recipient address is invalid")
	ErrMissingSubject   = errors.New("email: subject is required")
)

func validateEmailState(from string, to []string, subject string) error {
	if err := validateFrom(from); err != nil {
		return err
	}
	if err := validateRecipients(to); err != nil {
		return err
	}
	return validateSubject(subject)
}

func validateFrom(from string) error {
	if from == "" {
		return ErrMissingFrom
	}
	if !isValidAddress(from) {
		return ErrInvalidFrom
	}
	return nil
}

func validateRecipients(to []string) error {
	if len(to) == 0 {
		return ErrMissingRecipient
	}
	for _, address := range to {
		if !isValidAddress(address) {
			return ErrInvalidRecipient
		}
	}
	return nil
}

func validateSubject(subject string) error {
	if subject == "" {
		return ErrMissingSubject
	}
	return nil
}

func isValidAddress(address string) bool {
	if address == "" {
		return false
	}
	if !strings.Contains(address, atSymbol) {
		return false
	}
	if strings.HasPrefix(address, atSymbol) || strings.HasSuffix(address, atSymbol) {
		return false
	}
	return true
}
