package config

var (
	userAgent      string
	groupName      string
	subscriptionID string
)

//SetGroupName sets Group name
func SetGroupName(name string) {
	groupName = name
}

//GetGroupName return Group Name
func GroupName() string {
	return groupName
}

//SetSubscriptionName gets subscription Name
func SubscriptionID() string {
	return subscriptionID
}

// UserAgent specifies a string to append to the agent identifier.
func UserAgent() string {
	if len(userAgent) > 0 {
		return userAgent
	}
	return "radler-learnig-user-agent"
}
