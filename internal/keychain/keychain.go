// Package keychain wraps the go-keyring library
package keychain

import (
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

func Set(service, username, data string) error {
	return keyring.Set(service, username, data)
}

func Get(service, user string) (string, error) {
	data, err := keyring.Get(service, user)
	if err != nil {
		switch {
		case errors.Is(err, keyring.ErrNotFound):
			return "", fmt.Errorf("no data found for %s/%s", service, user)
		default:
			return "", fmt.Errorf("failed to get data from keychain: %w", err)
		}
	}
	return data, err
}

func Delete(service, user string) error {
	err := keyring.Delete(service, user)
	if err != nil {
		switch {
		case errors.Is(err, keyring.ErrNotFound):
			return fmt.Errorf("no data found for %s/%s", service, user)
		default:
			return fmt.Errorf("failed to delete from keychain: %w", err)
		}
	}
	return nil
}
