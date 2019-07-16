package planetlab

import (
	"testing"

	"gotest.tools/assert"
)

const existingSliceID = 1645
const nonExistingSliceID = -1
const existingSliceName = "ple_test"
const nonExistingSliceName = "foobarbaz"

func TestGetSlices(t *testing.T) {
}

func TestGetSliceByID(t *testing.T) {
	auth := GetTestAuth()

	// should return node with matching ID
	slice, err := GetSliceByID(auth, existingSliceID)
	assert.NilError(t, err)
	assert.Assert(t, slice.SliceID == existingSliceID)

	// should return err, since node with id -1 does not exist
	slice, err = GetSliceByID(auth, nonExistingSliceID)
	assert.Assert(t, err != nil)
	assert.Assert(t, slice == nil)
}

func TestGetSliceByName(t *testing.T) {
	auth := GetTestAuth()

	// should return slice with matching name
	slice, err := GetSliceByName(auth, existingSliceName)
	assert.NilError(t, err)
	assert.Assert(t, slice.Name == existingSliceName)

	// should return err, since node with hostname foo.bar.baz does not exist
	slice, err = GetSliceByName(auth, nonExistingSliceName)
	assert.Assert(t, err != nil)
	assert.Assert(t, slice == nil)
}

func TestUpdateSlice(t *testing.T) {

}
