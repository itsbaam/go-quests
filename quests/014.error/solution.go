package error

import (
	"errors"
	"fmt"
	"strings"
)

// Part 1: Define sentinel errors
var (
	ErrEmptyFilename = errors.New("filename cannot be empty")
	ErrFileTooLarge  = errors.New("file size exceeds limit")
)

// Part 2: Define custom error type
type ValidationError struct {
	Filename string
	Reason   string
}

// TODO: Implement Error() method
func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for '%s': %s", ve.Filename, ve.Reason)
}

// Part 3: Validation functions

func ValidateFilename(filename string) error {
	if filename == "" {
		return ErrEmptyFilename
	}
	return nil
}

func ValidateFileSize(size int64, maxSize int64) error {
	if size > maxSize {
		return ErrFileTooLarge
	} else if size < 0 {
		return errors.New("file size cannot be negative")
	}
	return nil
}

func ValidateFileExtension(filename string, allowedExts []string) error {
	for _, ext := range allowedExts {
		if strings.HasSuffix(filename, ext) {
			return nil
		}
	}
	return &ValidationError{
		Filename: filename,
		Reason:   "unsupported file extension",
	}
}

func ValidateFile(filename string, size int64, maxSize int64) error {
	if errName := ValidateFilename(filename); errName != nil {
		return fmt.Errorf("file validation failed: %w", errName)
	}
	if errSize := ValidateFileSize(size, maxSize); errSize != nil {
		return fmt.Errorf("file validation failed: %w", errSize)
	}
	return nil
}

// Part 4: Error checking

func CanRetry(err error) bool {
	if err == nil {
		return false
	}
	var ve *ValidationError
	return errors.As(err, &ve)
}
