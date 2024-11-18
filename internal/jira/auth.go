package jira

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/joakimen/ww/pkg/keychain"
	"golang.org/x/term"
)

const (
	serviceName = "ww-jira"
	username    = "jira-credentials"
)

type Credentials struct {
	Email    string `json:"email"`
	APIToken string `json:"api_token"`
	Domain   string `json:"domain"`
}

// Credential management functions
func SaveCredentials(email, apiToken, domain string) error {
	creds := Credentials{
		Email:    email,
		APIToken: apiToken,
		Domain:   domain,
	}

	err := keychain.Set(serviceName, username, creds)
	if err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	return nil
}

func GetCredentials() (*Credentials, error) {
	var creds Credentials
	err := keychain.Get(serviceName, username, &creds)
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}

	return &creds, nil
}

func DeleteCredentials() error {
	err := keychain.Delete(serviceName, username)
	if err != nil {
		return fmt.Errorf("failed to delete credentials: %w", err)
	}

	return nil
}

// Interactive command functions
func Login() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read email: %w", err)
	}
	email = strings.TrimSpace(email)

	fmt.Print("API Token: ")
	tokenBytes, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read token: %w", err)
	}
	fmt.Println()
	token := string(tokenBytes)

	fmt.Print("Domain: ")
	domain, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read domain: %w", err)
	}
	domain = strings.TrimSpace(domain)

	err = SaveCredentials(email, token, domain)
	if err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}

	fmt.Println("Successfully logged in to Jira")
	return nil
}

func Logout() error {
	err := DeleteCredentials()
	if err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}
	fmt.Println("Successfully logged out from Jira")
	return nil
}

func verifyCredentials(creds *Credentials) error {
	client := &http.Client{}
	url := fmt.Sprintf("https://%s/rest/api/3/myself", creds.Domain)

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(creds.Email, creds.APIToken)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid credentials")
	}

	var user struct {
		DisplayName  string `json:"displayName"`
		EmailAddress string `json:"emailAddress"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

func Status() error {
	creds, err := GetCredentials()
	if err != nil {
		fmt.Println("Not logged in to Jira. Use 'ww jira auth login' to authenticate.")
		return nil
	}

	err = verifyCredentials(creds)
	if err != nil {
		fmt.Printf("⚠️  Your Jira credentials are invalid or expired.\n")
		fmt.Printf("Please login again using 'ww jira auth login'\n")
		return nil
	}

	fmt.Printf("✓ Successfully authenticated to %s as %s\n", creds.Domain, creds.Email)
	return nil
}

func Show() error {
	creds, err := GetCredentials()
	if err != nil {
		return fmt.Errorf("failed to get credentials: %w", err)
	}
	fmt.Printf("Email:  %s\n", creds.Email)
	fmt.Printf("Domain: %s\n", creds.Domain)
	fmt.Printf("Token:  %s\n", maskToken(creds.APIToken))
	return nil
}

func maskToken(token string) string {
	if len(token) <= 8 {
		return "********"
	}
	return token[:4] + "..." + token[len(token)-4:]
}
