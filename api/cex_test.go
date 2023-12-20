package api_test

import (
	"testing"

	"github.com/devinmiller/fem-basics-of-go-client/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("Error was not found")
	}
}
