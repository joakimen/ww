// Package json provides helper functions for serializing and deserializing JSON data.
package json

import (
	"encoding/json"
	"fmt"
)

func Serialize[T any](data T) (string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to serialize data: %w", err)
	}
	return string(jsonBytes), nil
}

func Deserialize[T any](jsonStr string, target *T) error {
	if err := json.Unmarshal([]byte(jsonStr), target); err != nil {
		return fmt.Errorf("failed to deserialize data: %w", err)
	}
	return nil
}
