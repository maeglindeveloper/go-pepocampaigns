package pepo

import (
	"context"
	"net/url"

	"github.com/maeglindeveloper/go-pepocampaigns/domain"
)

// CreateListResponse defines the API response when creating a list
type CreateListResponse struct {
	BaseResponse
	Total uint64 `json:"total"`
	Data  struct {
		List domain.List `json:"list"`
	} `json:"data"`
}

// CreateList creates the contact list 'name' with specified parameters
// https://know.pepocampaigns.com/articles/managing-lists-api/
func (c *Client) CreateList(name string, source string, optInType *domain.OptInType, fromName, fromEmail *string) (*CreateListResponse, error) {
	params := &url.Values{}
	params.Add("name", name)

	params.Add("source", source)
	if optInType != nil {
		params.Add("opt_in_type", string(*optInType))
	}
	if fromName != nil {
		params.Add("from_name", *fromName)
	}
	if fromEmail != nil {
		params.Add("from_email", *fromEmail)
	}
	resp := &CreateListResponse{}
	if err := c.call(context.Background(), "POST", "list/create", params, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
