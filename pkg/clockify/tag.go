package clockify

import (
	"fmt"
	"net/url"
	"strconv"
)

type TagService struct {
	client *APIClient
}

type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
	Archived    bool   `json:"archived"`
}

type TagListFilter struct {
	Name     *string
	Archived *bool
	Page     int
	PageSize int
}

func (t *TagListFilter) ToQuery() string {
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

	return v.Encode()
}

func (s *TagService) List(workspaceID string, filter *TagListFilter) (*[]Tag, error) {
	path := fmt.Sprintf("workspaces/%s/tags", workspaceID)
	req, err := s.client.newAPIRequest("GET", path, filter.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	var tags []Tag
	_, err = s.client.do(req, &tags)
	if err != nil {
		return nil, err
	}

	return &tags, nil
}
