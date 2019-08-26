package planetlab

import (
	"testing"

	"gotest.tools/assert"
)

func TestAddPersonKey(t *testing.T) {
	auth := GetTestAuth()

	// add a dummy key
	keyID, err := AddPersonKey(auth, existingUser, "ssh", samplePubKeyContents)
	assert.NilError(t, err)
	assert.Assert(t, keyID > 0)

	// check dummy key exists
	keys, err := GetKeys(auth, struct {
		KeyID int `xmlrpc:"key_id"`
	}{keyID})

	assert.NilError(t, err)
	assert.Assert(t, len(keys) == 1)
	assert.Assert(t, keys[0].KeyID == keyID)
}

func TestDeleteKey(t *testing.T) {
	auth := GetTestAuth()

	// add a dummy key
	keyID, err := AddPersonKey(auth, existingUser, "ssh", samplePubKeyContents)
	assert.NilError(t, err)
	assert.Assert(t, keyID > 0)

	// delete dummy key
	err = DeleteKey(auth, keyID)
	assert.NilError(t, err)
}

func TestGetKeys(t *testing.T) {
	auth := GetTestAuth()

	// should return at least one key
	keys, err := GetKeys(auth, nil)
	assert.NilError(t, err)
	assert.Assert(t, len(keys) >= 1)
}
