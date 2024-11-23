// Package term wraps the golang.org/x/term package and provides
// helper functions for reading strings and passwords from the terminal.
package term

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func ReadString(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read email: %w", err)
	}
	return strings.TrimSpace(str), nil
}

func ReadPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	tokenBytes, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read token: %w", err)
	}
	fmt.Println()
	return string(tokenBytes), nil
}
