package jira

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
	email, err := term.ReadString("Jira Email: ")
	if err != nil {
		return fmt.Errorf("failed to read email: %w", err)
	}

	apiToken, err := term.ReadPassword("Jira API Token: ")
	if err != nil {
		return fmt.Errorf("failed to read API token: %w", err)
	}

	domain, err := term.ReadString("Jira Domain (e.g., mycompany.atlassian.net): ")
	if err != nil {
		return fmt.Errorf("failed to read domain: %w", err)
	}

	// Load existing creds or create new one
	creds, err := m.credentialsManager.Load()
	if err != nil {
		if !errors.Is(err, credentials.ErrKeychainCredentialsNotFound) {
			return fmt.Errorf("failed to load credentials: %w", err)
		}
		creds = credentials.NewCredentials()
	}

	// Add new account
	creds.Accounts.Jira = append(creds.Accounts.Jira, credentials.JiraAccount{
		Email:    email,
		APIToken: apiToken,
		Domain:   domain,
	})

	if err := m.credentialsManager.Save(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Println("Successfully logged in to Jira")
	return nil
}

func (m *CredentialsManager) Logout() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		return nil // Already logged out
	}

	creds.Accounts.Jira = nil // Remove all Jira accounts
	if err := m.credentialsManager.Save(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Println("Successfully logged out from Jira")
	return nil
}

func (m *CredentialsManager) Show() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		fmt.Println("Not logged in to Jira")
		return nil
	}

	if len(creds.Accounts.Jira) == 0 {
		fmt.Println("No Jira accounts configured")
		return nil
	}

	maskedCreds := credentials.MaskCredentials(creds)

	for i, acc := range maskedCreds.Accounts.Jira {
		if i > 0 {
			fmt.Println("---")
		}
		fmt.Printf("Email:  %s\n", credentials.MaskValue(acc.Email))
		fmt.Printf("Domain: %s\n", acc.Domain)
	}
	return nil
}

func (m *CredentialsManager) Status() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		fmt.Println("Not logged in to Jira")
		return nil
	}

	if len(creds.Accounts.Jira) == 0 {
		fmt.Println("Not logged in to Jira")
		return nil
	}

	fmt.Println("Logged in to Jira")
	return nil
}
