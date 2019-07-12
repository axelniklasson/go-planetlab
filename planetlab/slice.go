package planetlab

import (
	"fmt"
)

// Slice models a PlanetLab slice
type Slice struct {
	Creator           int    `xmlrpc:"creator_person_id"`
	Instantiation     string `xmlrpc:"instantiation"`
	SliceAttributeIDs []int  `xmlrpc:"slice_attribute_ids"`
	Name              string `xmlrpc:"name"`
	SliceID           int    `xmlrpc:"slice_id"`
	Created           int    `xmlrpc:"created"`
	URL               string `xmlrpc:"url"`
	MaxNodes          int    `xmlrpc:"max_nodes"`
	PersonIDs         []int  `xmlrpc:"person_ids"`
	Expires           int    `xmlrpc:"expires"`
	SiteID            int    `xmlrpc:"site_id"`
	PeerSliceID       int    `xmlrpc:"peer_slice_id"`
	NodeIDs           []int  `xmlrpc:"node_ids"`
	PeerID            int    `xmlrpc:"peer_id"`
	Description       string `xmlrpc:"description"`
}

func (s Slice) String() string {
	return fmt.Sprintf("Slice - SliceID: %d, SiteID: %d, NodeIDs: %v, URL: %s", s.SliceID, s.SiteID, s.NodeIDs, s.URL)
}

// API methods
const getSlicesMethod = "GetSlices"
const updateSliceMethod = "UpdateSlice"

// GetSlices returns a slice of Slices containing details about slices
func GetSlices(auth Auth, sliceFilter interface{}) ([]Slice, error) {
	client := GetClient()
	args := make([]interface{}, 2)
	args[0] = auth
	args[1] = sliceFilter
	slices := []Slice{}

	err := client.Call(getSlicesMethod, args, &slices)
	if err != nil {
		return nil, err
	}

	return slices, nil
}

// UpdateSlice updates the parameters of an existing slice
func UpdateSlice(auth Auth, s Slice, sliceFields interface{}) error {
	client := GetClient()
	args := make([]interface{}, 3)
	args[0] = auth
	args[1] = s.SliceID
	args[2] = sliceFields

	var result int
	err := client.Call(updateSliceMethod, args, &result)
	if err != nil {
		return err
	} else if result != 1 {
		return fmt.Errorf("Could not update slice. Result from API: %v", result)
	}

	return nil
}
