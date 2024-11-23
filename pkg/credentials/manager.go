// Package credentials provides a manager for storing and retrieving credentials in the keychain.
package credentials

import (
	"errors"
	"fmt"
	"strings"

	"github.com/joakimen/ww/internal/json"
	"github.com/joakimen/ww/internal/keychain"
)

const (
	KeychainServiceName = "ww"
	KeychainUsername    = "credentials"
)

var (
	ErrKeychainCredentialsNotFound     = errors.New("credentials not found in keychain")
	ErrKeychainCredentialsSaveFailed   = errors.New("failed to save credentials to keychain")
	ErrKeychainCredentialsLoadFailed   = errors.New("failed to load credentials from keychain")
	ErrKeychainCredentialsDeleteFailed = errors.New("failed to delete credentials from keychain")
	ErrKeychainCredentialsShowFailed   = errors.New("failed to show credentials from keychain")
)

type KeychainCredentialsManager struct{}

func NewKeychainCredentialsManager() *KeychainCredentialsManager {
	return &KeychainCredentialsManager{}
}

type Manager interface {
	Save(Credentials) error
	Load() (Credentials, error)
	Delete() error
	Show() error
}

func (m *KeychainCredentialsManager) Save(creds Credentials) error {
	jsonStr, err := json.Serialize(creds)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrKeychainCredentialsSaveFailed, err)
	}

	err = keychain.Set(KeychainServiceName, KeychainUsername, jsonStr)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrKeychainCredentialsSaveFailed, err)
	}
	return nil
}

func (m *KeychainCredentialsManager) Load() (Credentials, error) {
	creds := NewCredentials()

	jsonStr, err := keychain.Get(KeychainServiceName, KeychainUsername)
	if err != nil {
		return Credentials{}, fmt.Errorf("%w: %w", ErrKeychainCredentialsNotFound, err)
	}

	err = json.Deserialize(jsonStr, &creds)
	if err != nil {
		return Credentials{}, fmt.Errorf("%w: %w", ErrKeychainCredentialsLoadFailed, err)
	}

	return creds, nil
}

func (m *KeychainCredentialsManager) Delete() error {
	err := keychain.Delete(KeychainServiceName, KeychainUsername)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrKeychainCredentialsDeleteFailed, err)
	}
	return nil
}

func (m *KeychainCredentialsManager) Show() error {
	credsJSON, err := keychain.Get(KeychainServiceName, KeychainUsername)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrKeychainCredentialsShowFailed, err)
	}

	var creds Credentials
	err = json.Deserialize(credsJSON, &creds)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrKeychainCredentialsShowFailed, err)
	}

	credsMasked := MaskCredentials(creds)

	maskedCredsJSON, err := json.Serialize(credsMasked)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrKeychainCredentialsShowFailed, err)
	}

	fmt.Println(maskedCredsJSON)
	return nil
}

func MaskCredentials(creds Credentials) Credentials {
	maskedCreds := DeepCopy(creds)
	for i := range maskedCreds.Accounts.GitHub {
		maskedCreds.Accounts.GitHub[i].Token = MaskValue(maskedCreds.Accounts.GitHub[i].Token)
	}

	for i := range maskedCreds.Accounts.AWS {
		maskedCreds.Accounts.AWS[i].AccessKeyID = MaskValue(maskedCreds.Accounts.AWS[i].AccessKeyID)
		maskedCreds.Accounts.AWS[i].SecretAccessKey = MaskValue(maskedCreds.Accounts.AWS[i].SecretAccessKey)
	}

	for i := range maskedCreds.Accounts.Jira {
		maskedCreds.Accounts.Jira[i].APIToken = MaskValue(maskedCreds.Accounts.Jira[i].APIToken)
	}
	return maskedCreds
}

func MaskValue(value string) string {
	if len(value) <= 8 {
		return strings.Repeat("*", len(value))
	}
	return value[:4] + strings.Repeat("*", len(value)-8) + value[len(value)-4:]
}
