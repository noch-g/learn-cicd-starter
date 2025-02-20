package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}

	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected GetApiKey to fail on empty header")
	}

	header.Set("Authorization", "ApiK dsds")

	_, err = GetAPIKey(header)
	if err == nil {
		t.Fatal("expected GetApiKey to fail on wrong header")
	}

	header.Set("Authorization", "ApiKey cascdacaca")

	key, err := GetAPIKey(header)
	if err != nil {
		t.Fatal("GetApiKey failed on normally formated header")
	}
	if key != "cascdacaca" {
		t.Fatal("wrong key detected")
	}
}
