package planetlab

import "fmt"

// Key represents a PlanetLab key
type Key struct {
	PeerID    int    `xmlrpc:"peer_id"`
	KeyType   string `xmlrpc:"key_type"`
	Value     string `xmlrpc:"key"`
	PersonID  int    `xmlrpc:"person_id"`
	KeyID     int    `xmlrpc:"key_id"`
	PeerKeyID int    `xmlrpc:"peer_key_id"`
}

func (k Key) String() string {
	return fmt.Sprintf("Key - KeyID: %d, KeyType: %s, PersonID: %d, Value: %s", k.KeyID, k.KeyType, k.PersonID, k.Value)
}

// AddPersonKey adds a new key to the specified account
//
// URL: https://www.planet-lab.org/doc/plc_api#AddPersonKey
func AddPersonKey(auth Auth, email string, keyType string, keyValue string) (int, error) {
	client := GetClient()
	keyFields := struct {
		KeyType string `xmlrpc:"key_type"`
		Key     string `xmlrpc:"key"`
	}{keyType, keyValue}

	args := make([]interface{}, 3)
	args[0] = auth
	args[1] = email
	args[2] = keyFields

	var result int
	err := client.Call(addPersonKeyMethod, args, &result)
	if err != nil {
		return -1, err
	}

	if result < 0 {
		return -1, fmt.Errorf("Could not add key. Got error %d", result)
	}

	return result, nil
}

// DeleteKey deletes a given key
//
// URL: https://www.planet-lab.org/doc/plc_api#DeleteKey
func DeleteKey(auth Auth, keyID int) error {
	client := GetClient()
	args := make([]interface{}, 2)
	args[0] = auth
	args[1] = keyID

	var result int
	err := client.Call(deleteKeyMethod, args, &result)
	if err != nil {
		return err
	}

	if result != 1 {
		return fmt.Errorf("Could not delete key with ID %d. Got return code %d", keyID, result)
	}

	return nil
}

// GetKeys returns a slice of keys
//
// URL: https://www.planet-lab.org/doc/plc_api#GetKeys
func GetKeys(auth Auth, keyFilter interface{}) ([]Key, error) {
	client := GetClient()
	args := make([]interface{}, 2)
	args[0] = auth
	args[1] = keyFilter
	keys := []Key{}

	err := client.Call(getKeysMethod, args, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}
