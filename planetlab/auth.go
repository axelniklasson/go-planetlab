package planetlab

import (
	"log"

	"github.com/kolo/xmlrpc"
)

var clientInstance *xmlrpc.Client

const plAPIURL = "https://www.planet-lab.eu/PLCAPI/"

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

// GetClient creates and returns
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
