package planetlab

// Key represents a PlanetLab key
type Key struct {
}

func (k Key) String() string {
	return ""
}

// AddPersonKey adds a new key to the specified account
func AddPersonKey() (int, error) {
	return 1, nil
}

// DeleteKey deletes a key (..)
func DeleteKey() error {
	return nil
}

// GetKeys returns a slice of keys
func GetKeys() ([]Key, error) {
	return []Key{}, nil
}

// UpdateKey pdates the parameters of an existing key
func UpdateKey() error {
	return nil
}
