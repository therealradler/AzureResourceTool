package config

import "os"

//Set Environmental Variables
func ParseEnvironment() error {

	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	return nil

}
