package github

import (
	"errors"
	"fmt"

	"github.com/joakimen/ww/internal/term"
	"github.com/joakimen/ww/pkg/credentials"
)

type CredentialsManager struct {
	credentialsManager credentials.Manager
}

func NewCredentialsManager(credentialsManager credentials.Manager) *CredentialsManager {
	return &CredentialsManager{credentialsManager}
}

func (m *CredentialsManager) Login() error {
	token, err := term.ReadPassword("GitHub Token: ")
	if err != nil {
		return fmt.Errorf("failed to read token: %w", err)
	}

	creds, err := m.credentialsManager.Load()
	if err != nil {
		if !errors.Is(err, credentials.ErrKeychainCredentialsNotFound) {
			return fmt.Errorf("failed to load credentials: %w", err)
		}
		creds = credentials.NewCredentials()
	}

	fmt.Printf("store: %+v\n", creds)

	creds.Accounts.GitHub = append(creds.Accounts.GitHub, credentials.GitHubAccount{Token: token})
	if err := m.credentialsManager.Save(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Println("Successfully logged in to GitHub")
	return nil
}

func (m *CredentialsManager) Logout() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		return nil // Already logged out
	}

	creds.Accounts.GitHub = nil // Remove all GitHub accounts
	if err := m.credentialsManager.Save(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Println("Successfully logged out from GitHub")
	return nil
}

func (m *CredentialsManager) Show() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		fmt.Println("Not logged in to GitHub")
		return nil
	}

	if len(creds.Accounts.GitHub) == 0 {
		fmt.Println("No GitHub accounts configured")
		return nil
	}

	maskedCreds := credentials.MaskCredentials(creds)
	for i, acc := range maskedCreds.Accounts.GitHub {
		if i > 0 {
			fmt.Println("---")
		}
		fmt.Printf("Token: %s\n", credentials.MaskValue(acc.Token))
	}
	return nil
}

func (m *CredentialsManager) Status() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		fmt.Println("Not logged in to GitHub")
		return nil
	}

	if len(creds.Accounts.GitHub) == 0 {
		fmt.Println("Not logged in to GitHub")
		return nil
	}

	fmt.Println("Logged in to GitHub")
	return nil
}
