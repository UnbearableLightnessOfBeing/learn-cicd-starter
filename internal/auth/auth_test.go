package auth_test

import (
	"errors"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestFailsIfAuthHeaderNotIncluded(t *testing.T) {
	header := make(map[string][]string)

	key, err := auth.GetAPIKey(header)
	if key != "" {
		t.Fatal("key should be empty")
	}
	if !errors.Is(err, auth.ErrNoAuthHeaderIncluded) {
		t.Fatal("should be header not included error")
	}
}

func TestFailsOnMalformedAuthHeader(t *testing.T) {
	header := make(map[string][]string)
	header["Authorization"] = []string{"ApiKey_aboba"}

	key, err := auth.GetAPIKey(header)
	if key != "" {
		t.Fatal("key should be empty")
	}
	if err == nil {
		t.Fatal("should error")
	}
}

func TestGetApiKeySuccessfully(t *testing.T) {
	header := make(map[string][]string)
	header["Authorization"] = []string{"ApiKey aboba"}

	key, err := auth.GetAPIKey(header)
	if err != nil {
		t.Fatal("shouldn't error")
	}
	if key != "aboba" {
		t.Fatal("go wrong key")
	}
}
