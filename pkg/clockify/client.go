package clockify

import (
	"fmt"
	"net/url"
	"strconv"
)

// ClientService serve the clockify client api
type ClientService struct {
	client *APIClient
}

// Client resource from clockify
type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
	Archived    bool   `json:"archived"`
}

// ClientListFilter for the clockify list endpoint
type ClientListFilter struct {
	Name     *string
	Archived *bool
	Page     int
	PageSize int
}

// ToQuery formats the filters for the Get Query
func (c *ClientListFilter) ToQuery() string {
	v := url.Values{}

	if c.Name != nil {
		v.Add("name", *c.Name)
	}
	if c.Archived != nil {
		v.Add("archived", strconv.FormatBool(*c.Archived))
	}

	if c.Page > 0 {
		v.Add("page", strconv.Itoa(c.Page))
	}
	if c.PageSize > 0 {
		v.Add("page-size", strconv.Itoa(c.PageSize))
	}

	return v.Encode()
}

// List all clients in a workspace
func (s *ClientService) List(workspaceID string, filter *ClientListFilter) (*[]Client, error) {
	path := fmt.Sprintf("workspaces/%s/clients", workspaceID)
	req, err := s.client.newAPIRequest("GET", path, filter.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	var clients []Client
	_, err = s.client.do(req, &clients)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}
