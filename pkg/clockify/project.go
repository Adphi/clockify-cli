package clockify

import (
	"fmt"
	"net/url"
	"strconv"
)

type ProjectService struct {
	client *APIClient
}

type ProjectEstimate struct {
	Estimate string `json:"estimate"`
	Type     string `json:"type"`
}

type Project struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	ClientID    string          `json:"clientId"`
	WorkspaceID string          `json:"workspaceId"`
	Billable    bool            `json:"billable"`
	Memberships []Membership    `json:"memberships"`
	Color       string          `json:"color"`
	Estimate    ProjectEstimate `json:"estimate"`
	Archived    bool            `json:"archived"`
	Duration    string          `json:"duration"`
	ClientName  string          `json:"clientName"`
	Note        string          `json:"note"`
	Template    bool            `json:"template"`
	Public      bool            `json:"public"`
}

type ProjectListFilter struct {
	Name            *string
	Archived        *bool
	Page            int
	PageSize        int
	Billable        *bool
	ClientIDs       *[]string
	ContainsClients *bool
	UserIDs         *[]string
	ContainsUsers   *bool
	IsTemplate      *bool
	SortColum       *string
	SortOrder       *string
}

func (t *ProjectListFilter) ToQuery() string {
	v := url.Values{}

	if t.Name != nil {
		v.Add("name", *t.Name)
	}
	if t.Archived != nil {
		v.Add("archived", strconv.FormatBool(*t.Archived))
	}

	if t.Page > 0 {
		v.Add("page", strconv.Itoa(t.Page))
	}
	if t.PageSize > 0 {
		v.Add("page-size", strconv.Itoa(t.PageSize))
	}

	if t.Billable != nil {
		v.Add("billable", strconv.FormatBool(*t.Billable))
	}

	if t.ClientIDs != nil {
		for _, clientID := range *t.ClientIDs {
			v.Add("clients", clientID)
		}
	}
	if t.ContainsClients != nil {
		v.Add("contains-client", strconv.FormatBool(*t.ContainsClients))
	}

	if t.UserIDs != nil {
		for _, userID := range *t.UserIDs {
			v.Add("users", userID)
		}
	}
	if t.ContainsUsers != nil {
		v.Add("contains-users", strconv.FormatBool(*t.ContainsUsers))
	}

	if t.IsTemplate != nil {
		v.Add("is-template", strconv.FormatBool(*t.IsTemplate))
	}

	if t.SortColum != nil {
		v.Add("sort-colum", *t.SortColum)
	}
	if t.SortOrder != nil {
		v.Add("sort-order", *t.SortOrder)
	}

	return v.Encode()
}

func (s *ProjectService) List(workspaceID string, filter *ProjectListFilter) (*[]Project, error) {
	path := fmt.Sprintf("workspaces/%s/projects", workspaceID)
	req, err := s.client.newAPIRequest("GET", path, filter.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	var projects []Project
	_, err = s.client.do(req, &projects)
	if err != nil {
		return nil, err
	}

	return &projects, nil
}

func (s *ProjectService) Get(workspaceID string, projectID string) (*Project, error) {
	path := fmt.Sprintf("workspaces/%s/projects/%s", workspaceID, projectID)
	req, err := s.client.newAPIRequest("GET", path, "", nil)
	if err != nil {
		return nil, err
	}

	var project Project
	_, err = s.client.do(req, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
