package config

import (
	"errors"
	"os"
)

func init() {
	os.Setenv("AZURE_SUBSCRIPTION_ID", "f135b611-b863-4f6a-a5ec-507b1111641b")
}

//Set Environmental Variables
func ParseEnvironment() error {
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		return errors.New("Sub Id Empty")
	}
	return nil

}
