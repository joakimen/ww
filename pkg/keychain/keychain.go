package keychain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

// Set saves data to the system keychain
func Set(serviceName, username string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = keyring.Set(serviceName, username, string(jsonData))
	if err != nil {
		return fmt.Errorf("failed to save to keychain: %w", err)
	}

	return nil
}

// Get retrieves data from the system keychain
func Get(serviceName, username string, target interface{}) error {
	data, err := keyring.Get(serviceName, username)
	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return fmt.Errorf("no data found for %s/%s", serviceName, username)
		}
		return fmt.Errorf("failed to get data from keychain: %w", err)
	}

	if err := json.Unmarshal([]byte(data), target); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return nil
}

// Delete removes data from the system keychain
func Delete(serviceName, username string) error {
	err := keyring.Delete(serviceName, username)
	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return fmt.Errorf("no data found for %s/%s", serviceName, username)
		}
		return fmt.Errorf("failed to delete from keychain: %w", err)
	}

	return nil
}
