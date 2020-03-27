package main

import (
	"context"
	"fmt"
	"log"

	"github.com/therealradler/AzureResourceTool/config"

	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
)

type clientInfo struct {
	SubscriptionID string
}

var (
	ctx        = context.Background()
	clientData clientInfo
	authorizer autorest.Authorizer
)

func init() {
	var err error
	authorizer, err = auth.NewAuthorizerFromCLI()
	if err != nil {
		log.Fatalf("Failed to Authetnicate From Environment: %v", err)
	}
	config.ParseEnvironment()
	clientData.SubscriptionID = config.SubscriptionID()

}

func main() {
	var err error
	fmt.Print(config.SubscriptionID())
	resourcesClient := resources.NewClient(clientData.SubscriptionID)
	resourcesClient.Authorizer = authorizer
	config.SetGroupName("Versacomm-qa")
	results, err := GetGroups(ctx)
	if err != nil {
		fmt.Printf("Failed to Return Group %v", err)
	}
	fmt.Println(*results.Name)
	fmt.Println(*results.Location)
	fmt.Println(*results.ID)

}

func getClient() resources.GroupsClient {
	groupClient := resources.NewGroupsClient(clientData.SubscriptionID)
	groupClient.Authorizer = authorizer
	groupClient.AddToUserAgent(config.UserAgent())
	return groupClient
}

func GetGroups(ctx context.Context) (resources.Group, error) {
	groupsClient := getClient()
	return groupsClient.Get(ctx, config.GroupName())
}
