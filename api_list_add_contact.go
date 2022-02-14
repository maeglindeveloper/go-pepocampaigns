package pepo

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/maeglindeveloper/go-pepocampaigns/domain"
)

// UserStatusesInput defines the user status input (contact)
type UserStatusesInput struct {
	DoubleOptInStatus *domain.DoubleOptInStatus
	SubscribeStatus   *domain.SubscribeStatus
	BlacklistStatus   *domain.BlacklistStatus
}

// AddContactOptions defines the options structure when adding a contact
type AddContactOptions struct {
	Attributes   *map[string]interface{}
	UserStatuses *UserStatusesInput
}

// AddContactResponse defines the API response when adding a contact to a list
// https://know.pepocampaigns.com/articles/managing-lists-api/
type AddContactResponse struct {
	BaseResponse
	Email string `json:"email"`
}

// AddContact add a contact email to the specified listID with options
func (c *Client) AddContact(listID uint64, email string, options *AddContactOptions) (*AddContactResponse, error) {
	params := &url.Values{}
	params.Add("list_id", strconv.FormatUint(listID, 10))
	params.Add("email", email)

	if options != nil {
		if options.Attributes != nil {
			for k, attr := range *options.Attributes {
				formattedKey := fmt.Sprintf("attributes[%s]", k)
				switch v := attr.(type) {
				case string:
					params.Add(formattedKey, v)
				case uint64:
					params.Add(formattedKey, strconv.FormatUint(v, 10))
				case int64:
					params.Add(formattedKey, strconv.FormatInt(v, 10))
				case time.Time:
					params.Add(formattedKey, v.Format(time.RFC3339))
				}
			}
		}
		if options.UserStatuses != nil {
			if options.UserStatuses.BlacklistStatus != nil {
				params.Add("user_status[blacklist_status]", string(*options.UserStatuses.BlacklistStatus))
			}
			if options.UserStatuses.SubscribeStatus != nil {
				params.Add("user_status[subscribe_status]", string(*options.UserStatuses.SubscribeStatus))
			}
			if options.UserStatuses.DoubleOptInStatus != nil {
				params.Add("user_status[double_opt_in_status]", string(*options.UserStatuses.DoubleOptInStatus))
			}
		}
	}

	resp := &AddContactResponse{}
	if err := c.call(context.Background(), "POST", fmt.Sprintf("list/%d/add-contact", listID), params, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
