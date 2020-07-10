package clockify

import (
	"fmt"
	"net/url"
	"strconv"
)

// TaskService serve the clockify task api
type TaskService struct {
	client *APIClient
}

// Task resource from clockify
type Task struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ProjectID string `json:"projectId"`

	Status      string   `json:"status"`
	Estimate    string   `json:"estimate"`
	AssigneeIds []string `json:"assigneeIds"`
}

// TaskListFilter for the clockify list endpoint
type TaskListFilter struct {
	Name     *string
	IsActive *bool
	Page     int
	PageSize int
}

// ToQuery formats the filters for the Get Query
func (t *TaskListFilter) ToQuery() string {
	v := url.Values{}

	if t.Name != nil {
		v.Add("name", *t.Name)
	}
	if t.IsActive != nil {
		v.Add("is-active", strconv.FormatBool(*t.IsActive))
	}

	if t.Page > 0 {
		v.Add("page", strconv.Itoa(t.Page))
	}
	if t.PageSize > 0 {
		v.Add("page-size", strconv.Itoa(t.PageSize))
	}

	return v.Encode()
}

// List all task in a workspace and project
func (t *TaskService) List(workspaceID string, projectID string, filter *TaskListFilter) (*[]Task, error) {
	path := fmt.Sprintf("workspaces/%s/projects/%s/tasks", workspaceID, projectID)
	req, err := t.client.newAPIRequest("GET", path, filter.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	_, err = t.client.do(req, &tasks)
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}

// Get a task by id in a workspace and project
func (t *TaskService) Get(workspaceID string, projectID string, taskID string) (*Task, error) {
	path := fmt.Sprintf("workspaces/%s/projects/%s/tasks/%s", workspaceID, projectID, taskID)
	req, err := t.client.newAPIRequest("GET", path, "", nil)
	if err != nil {
		return nil, err
	}

	var task Task
	_, err = t.client.do(req, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
