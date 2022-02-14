package domain

// SubscribeStatus defines the enumeration for the subscription status of a contact
type SubscribeStatus string

const (
	// SubscribeStatusSubscribed defines that the contact has subscribed to pepo
	SubscribeStatusSubscribed SubscribeStatus = "subscribed"
	// SubscribeStatusUnsubscribed defines that the contact has unsubscribed to pepo
	SubscribeStatusUnsubscribed SubscribeStatus = "unsubscribed"
)

// BlacklistStatus defines the enumration for blacklist status of a contact
type BlacklistStatus string

const (
	// BlacklistStatusBlacklisted defines that the contact is blacklisted
	BlacklistStatusBlacklisted BlacklistStatus = "blacklisted"
	// BlacklistStatusUnblacklisted defines that the contact is unblacklisted
	BlacklistStatusUnblacklisted BlacklistStatus = "unblacklisted"
)

// DoubleOptInStatus defines the status of the double option
type DoubleOptInStatus string

const (
	// DoubleOptInStatusVerified defines that the contact has approved to be listed on pepo (by confirming it)
	DoubleOptInStatusVerified DoubleOptInStatus = "verified"
	// DoubleOptInStatusPending defines that the contact is still pending to approve to be listed on pepo
	DoubleOptInStatusPending DoubleOptInStatus = "pending"
)
