package planetlab

// Key represents a PlanetLab key
type Key struct {
}

func (k Key) String() string {
	return ""
}

// AddPersonKey adds a new key to the specified account
//
// URL: https://www.planet-lab.org/doc/plc_api#AddPersonKey
func AddPersonKey() (int, error) {
	return 1, nil
}

// DeleteKey deletes a given key
//
// URL: https://www.planet-lab.org/doc/plc_api#DeleteKey
func DeleteKey() error {
	return nil
}

// GetKeys returns a slice of keys
//
// URL: https://www.planet-lab.org/doc/plc_api#GetKeys
func GetKeys() ([]Key, error) {
	return []Key{}, nil
}

// UpdateKey updates the parameters of an existing key
//
// URL: https://www.planet-lab.org/doc/plc_api#UpdateKey
func UpdateKey() error {
	return nil
}
