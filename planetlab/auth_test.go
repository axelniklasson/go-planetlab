package planetlab

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetClientAuth(t *testing.T) {
}

func TestGetClient(t *testing.T) {

}

func TestAuthenticate(t *testing.T) {
	// should succeed if env setup correctly
	auth := GetTestAuth()
	err := auth.Authenticate()
	assert.Assert(t, err == nil)

	// should fail since password is now empty
	auth.AuthString = ""
	err = auth.Authenticate()
	assert.Assert(t, err != nil)
}
