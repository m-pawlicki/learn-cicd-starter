package tests

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKeyEmptyHeader(t *testing.T) {
	headers := http.Header{}
	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Error("want error for invalid input")
	}

}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "123456")
	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Error("want error for invalid input")
	}
}

func TestGetAPIKeyWrongKindHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer 123456")
	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Error("want error for invalid input")
	}
}

func TestGetAPIKeyGoodHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 123456")
	expected := "123456"
	key, err := auth.GetAPIKey(headers)
	if err != nil {
		t.Error("want error for invalid input")
	}
	if key != expected {
		t.Errorf("got %s, expected %s", key, expected)
	}
}
