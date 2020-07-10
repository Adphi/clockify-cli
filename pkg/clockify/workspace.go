package clockify

import (
	"fmt"
	"time"
)

type WorkspaceService struct {
	client *APIClient
}

type Workspace struct {
	// todo: hourlyRate
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	ImageURL    string            `json:"imageUrl"`
	Memberships []Membership      `json:"memberships"`
	Settings    WorkspaceSettings `json:"workspaceSettings"`
}

type WorkspaceSettings struct {
	CanSeeTimeSheet                    bool                    `json:"canSeeTimeSheet"`
	CanSeeTracker                      bool                    `json:"canSeeTracker"`
	IsDefaultBillableProjects          bool                    `json:"defaultBillableProjects"`
	ForceDescription                   bool                    `json:"forceDescription"`
	ForceProjects                      bool                    `json:"forceProjects"`
	ForceTags                          bool                    `json:"forceTags"`
	ForceTasks                         bool                    `json:"forceTasks"`
	OnlyAdminsCreateProject            bool                    `json:"onlyAdminsCreateProject"`
	OnlyAdminsCreateTag                bool                    `json:"onlyAdminsCreateTag"`
	OnlyAdminsCreateTask               bool                    `json:"onlyAdminsCreateTask"`
	OnlyAdminsSeeAllTimeEntries        bool                    `json:"onlyAdminsSeeAllTimeEntries"`
	OnlyAdminsSeeBillableRates         bool                    `json:"onlyAdminsSeeBillableRates"`
	OnlyAdminsSeeDashboard             bool                    `json:"onlyAdminsSeeDashboard"`
	OnlyAdminsSeePublicProjectsEntries bool                    `json:"onlyAdminsSeePublicProjectsEntries"`
	ProjectFavorites                   bool                    `json:"projectFavorites"`
	ProjectGroupingLabel               string                  `json:"projectGroupingLabel"`
	ProjectPickerSpecialFilter         bool                    `json:"projectPickerSpecialFilter"`
	TimeRoundingInReports              bool                    `json:"timeRoundingInReports"`
	TrackTimeDownToSecond              bool                    `json:"trackTimeDownToSecond"`
	IsProjectPublicByDefault           bool                    `json:"isProjectPublicByDefault"`
	FeatureSubscriptionType            string                  `json:"featureSubscriptionType"`
	Round                              *WorkspaceRoundSettings `json:"round"`
	AutomaticLock                      *AutomaticLockSettings  `json:"automaticLock"`
	LockTimeEntries                    *time.Time              `json:"lockTimeEntries"`
}

type WorkspaceRoundSettings struct {
	Minutes string `json:"minutes"`
	Round   string `json:"round"`
}

type AutomaticLockSettings struct {
	ChangeDay       string `json:"changeDay"`
	DayOfMonth      int    `json:"dayOfMonth"`
	FirstDay        string `json:"firstDay"`
	OlderThanPeriod string `json:"olderThanPeriod"`
	OlderThanValue  int    `json:"olderThanValue"`
	Type            string `json:"type"`
}

func (s *WorkspaceService) List() (*[]Workspace, error) {
	path := fmt.Sprintf("workspaces")
	req, err := s.client.newAPIRequest("GET", path, "", nil)
	if err != nil {
		return nil, err
	}

	var workspaces []Workspace
	_, err = s.client.do(req, &workspaces)
	if err != nil {
		return nil, err
	}

	return &workspaces, nil
}
