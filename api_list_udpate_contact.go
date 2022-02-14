package pepo

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// UpdateContactOptions defines the update contact options
type UpdateContactOptions struct {
	Attributes   *map[string]interface{}
	UserStatuses *UserStatusesInput
}

// UpdateContactResponse defines the API response to an update of contact
type UpdateContactResponse struct {
	BaseResponse
	Email string `json:"email"`
}

// UpdateContact updates the contact 'email' from a 'listID' using options
func (c *Client) UpdateContact(listID uint64, email string, options *UpdateContactOptions) (*UpdateContactResponse, error) {
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

	resp := &UpdateContactResponse{}
	if err := c.call(context.Background(), "POST", fmt.Sprintf("list/%d/update-contact", listID), params, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
