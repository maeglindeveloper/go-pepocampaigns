package pepo

import (
	"context"
	"net/url"
)

// UpdateUserStatusType defines the entry to update user status
type UpdateUserStatusType string

const (
	UpdateUserStatusTypeUnsubscribe UpdateUserStatusType = "unsubscribe"
	UpdateUserStatusTypeBlacklist   UpdateUserStatusType = "blacklist"
	UpdateUserStatusTypeResubscribe UpdateUserStatusType = "resubscribe"
)

// UpdateUserStatusResponse defines the API response of an update user status
type UpdateUserStatusResponse struct {
	BaseResponse
	Email string `json:"email"`
}

// UpdateUserStatus updates the contact 'email' status
func (c *Client) UpdateUserStatus(email string, status UpdateUserStatusType) (*UpdateUserStatusResponse, error) {
	params := &url.Values{}
	params.Add("email", email)
	params.Add("type", string(status))
	resp := &UpdateUserStatusResponse{}
	if err := c.call(context.Background(), "POST", "user/change-status", params, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
