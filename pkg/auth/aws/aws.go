package aws

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
	accessKeyID, err := term.ReadString("AWS Access Key ID: ")
	if err != nil {
		return fmt.Errorf("failed to read access key ID: %w", err)
	}

	secretKey, err := term.ReadPassword("AWS Secret Access Key: ")
	if err != nil {
		return fmt.Errorf("failed to read secret access key: %w", err)
	}

	region, err := term.ReadString("AWS Region: ")
	if err != nil {
		return fmt.Errorf("failed to read region: %w", err)
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
	creds.Accounts.AWS = append(creds.Accounts.AWS, credentials.AWSAccount{
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretKey,
		Region:          region,
	})

	if err := m.credentialsManager.Save(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Println("Successfully logged in to AWS")
	return nil
}

func (m *CredentialsManager) Logout() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		return nil // Already logged out
	}

	creds.Accounts.AWS = nil // Remove all AWS accounts
	if err := m.credentialsManager.Save(creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Println("Successfully logged out from AWS")
	return nil
}

func (m *CredentialsManager) Show() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		fmt.Println("Not logged in to AWS")
		return nil
	}

	if len(creds.Accounts.AWS) == 0 {
		fmt.Println("No AWS accounts configured")
		return nil
	}

	maskedCreds := credentials.MaskCredentials(creds)
	for i, acc := range maskedCreds.Accounts.AWS {
		if i > 0 {
			fmt.Println("---")
		}

		fmt.Printf("Secret Key:    %s\n", credentials.MaskValue(acc.SecretAccessKey))
		fmt.Printf("Access Key ID: %s\n", acc.AccessKeyID)
		fmt.Printf("Region:        %s\n", acc.Region)
	}
	return nil
}

func (m *CredentialsManager) Status() error {
	creds, err := m.credentialsManager.Load()
	if err != nil {
		fmt.Println("Not logged in to AWS")
		return nil
	}

	if len(creds.Accounts.AWS) == 0 {
		fmt.Println("Not logged in to AWS")
		return nil
	}

	fmt.Println("Logged in to AWS")
	return nil
}
