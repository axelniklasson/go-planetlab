package planetlab

import (
	"log"
	"os"

	"github.com/kolo/xmlrpc"
)

var clientInstance *xmlrpc.Client

// Auth models the authentication object to use when authenticating against PL API
type Auth struct {
	AuthMethod string
	Username   string
	AuthString string
}

// GetClientAuth returns an Auth struct needed to authenticate against PL API
func GetClientAuth(username string, password string) Auth {
	return Auth{
		"password",
		username,
		password,
	}
}

// Authenticate can be used to test that the Auth object successfully can authenticate with the API
func (a Auth) Authenticate() error {
	client := GetClient()
	args := make([]interface{}, 1)
	args[0] = a
	var success int

	err := client.Call(authCheckMethod, args, &success)
	if err != nil {
		return err
	}

	return nil
}

// GetTestAuth returns an Auth object to be used during testing
func GetTestAuth() Auth {
	username := os.Getenv(plUsername)
	password := os.Getenv(plPassword)

	if username == "" || password == "" {
		log.Fatal("Could not load PL API credentials. Make sure both PL_USERNAME and PL_PASSWORD env variables are set.")
	}

	return GetClientAuth(username, password)
}

// GetClient creates and returns a client to use with the API
func GetClient() *xmlrpc.Client {
	if clientInstance == nil {
		clientInstance, err := xmlrpc.NewClient(plAPIURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		return clientInstance
	}
	return clientInstance
}
