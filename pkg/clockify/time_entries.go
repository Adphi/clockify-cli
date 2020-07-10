package clockify

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type TimeEntryService struct {
	client *APIClient
}

type TimeInterval struct {
	Duration string     `json:"duration"`
	Start    time.Time  `json:"start"`
	End      *time.Time `json:"end"`
}

type TimeEntry struct {
	ID          string   `json:"id"`
	WorkspaceID string   `json:"workspaceId"`
	UserID      string   `json:"userId"`
	ProjectID   *string  `json:"projectId"`
	TaskID      *string  `json:"taskId"`
	TagIDs      []string `json:"tagIds"`

	Description  string       `json:"description"`
	Billable     bool         `json:"billable"`
	IsLocked     bool         `json:"isLocked"`
	TimeInterval TimeInterval `json:"timeInterval"`
}

type TimeEntryListFilter struct {
	Start       *time.Time
	End         *time.Time
	Description *string
	ProjectID   *string
	TaksID      *string
	TagIDs      []string
	Page        int
	PageSize    int
}

func (t *TimeEntryListFilter) ToQuery() string {
	v := url.Values{}

	if t.Start != nil {
		v.Add("start", t.Start.UTC().Format(time.RFC3339))
	}
	if t.End != nil {
		v.Add("end", t.End.UTC().Format(time.RFC3339))
	}
	if t.Description != nil {
		v.Add("description", *t.Description)
	}
	if t.ProjectID != nil {
		v.Add("project", *t.ProjectID)
	}
	if t.TaksID != nil {
		v.Add("task", *t.TaksID)
	}

	for _, tagID := range t.TagIDs {
		v.Add("tags", tagID)
	}

	if t.Page > 0 {
		v.Add("page", strconv.Itoa(t.Page))
	}
	if t.PageSize > 0 {
		v.Add("page-size", strconv.Itoa(t.PageSize))
	}

	return v.Encode()
}

func (s *TimeEntryService) List(workspaceID string, userID string, filter *TimeEntryListFilter) (*[]TimeEntry, error) {
	path := fmt.Sprintf("workspaces/%s/user/%s/time-entries", workspaceID, userID)
	req, err := s.client.newAPIRequest("GET", path, filter.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	var timeEntries []TimeEntry
	_, err = s.client.do(req, &timeEntries)
	if err != nil {
		return nil, err
	}

	return &timeEntries, nil
}
