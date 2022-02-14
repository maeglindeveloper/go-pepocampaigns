package pepo

import (
	"context"
	"net/url"
	"strconv"

	"github.com/maeglindeveloper/go-pepocampaigns/domain"
)

// GetListsResponse defines the API response of a get all lists
type GetListsResponse struct {
	BaseResponse
	Total uint64        `json:"total"`
	Data  []domain.List `json:"data"`
}

// GetLists returns the pepocampaigns list specified by page
func (c *Client) GetLists(page uint64) (*GetListsResponse, error) {
	params := &url.Values{}
	params.Add("page", strconv.FormatUint(page, 10))
	resp := &GetListsResponse{}
	if err := c.call(context.Background(), "GET", "lists", params, nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
