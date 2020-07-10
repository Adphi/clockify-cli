package clockify

import (
	"fmt"
	"net/url"
	"strconv"
)

// UserService serve the clockify user api
type UserService struct {
	client *APIClient
}

type Membership struct {
	// todo: hourlyRate
	MembershipStatus string `json:"membershipStatus"`
	MembershipType   string `json:"membershipType"`
	TargetID         string `json:"targetId"`
	UserID           string `json:"userId"`
}

type UserSettings struct {
	TimeZone  string `json:"timeZone"`
	WeekStart string `json:"weekStart"`
}

// UserInfo resource from clockify
type UserInfo struct {
	ActiveWorkspace  string `json:"activeWorkspace"`
	DefaultWorkspace string `json:"defaultWorkspace"`

	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
	Status         string `json:"status"`

	Memberships []Membership `json:"memberships"`
	Settings    UserSettings `json:"settings"`
}

// Info about the current user
func (s *UserService) Info() (*UserInfo, error) {
	req, err := s.client.newAPIRequest("GET", "user", "", nil)
	if err != nil {
		return nil, err
	}

	var userInfo UserInfo
	_, err = s.client.do(req, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// UserListFilter for the clockify list endpoint
type UserListFilter struct {
	Name        *string
	Page        int
	PageSize    int
	Memberships *string
	Email       *string
	ProjectID   *string
	Status      *string

	SortColum *string
	SortOrder *string
}

// ToQuery formats the filters for the Get Query
func (u *UserListFilter) ToQuery() string {
	v := url.Values{}

	if u.Name != nil {
		v.Add("name", *u.Name)
	}

	if u.Page > 0 {
		v.Add("page", strconv.Itoa(u.Page))
	}
	if u.PageSize > 0 {
		v.Add("page-size", strconv.Itoa(u.PageSize))
	}

	if u.Email != nil {
		v.Add("email", *u.Email)
	}

	if u.ProjectID != nil {
		v.Add("projectId", *u.ProjectID)
	}

	if u.Status != nil {
		v.Add("status", *u.Status)
	}

	if u.Memberships != nil {
		v.Add("memberships", *u.Memberships)
	}

	if u.SortColum != nil {
		v.Add("sort-colum", *u.SortColum)
	}
	if u.SortOrder != nil {
		v.Add("sort-order", *u.SortOrder)
	}

	return v.Encode()
}

// List all users in a workspace
func (u *UserService) List(workspaceID string, filter *UserListFilter) (*[]UserInfo, error) {
	path := fmt.Sprintf("workspaces/%s/users", workspaceID)
	req, err := u.client.newAPIRequest("GET", path, filter.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	var users []UserInfo
	_, err = u.client.do(req, &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
