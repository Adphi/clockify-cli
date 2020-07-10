package clockify

import (
	"fmt"
	"time"
)

type ReportService struct {
	client *APIClient
}

type SummaryFilter struct {
	Groups       []string `json:"group"`      // (PROJECT, CLIENT, TASK, TAG, DATE, USER, USER_GROUP, TIMEENTRY)
	SortColumnBy string   `json:"sortColumn"` // (GROUP, DURATION, AMOUNT)
}

type ReportTagsFilter struct {
	IDs      []string `json:"ids"`
	Contains string   `json:"contains"` // (CONTAINS, DOES_NOT_CONTAIN, CONTAINS_ONLY)
	Status   string   `json:"status"`   // (ALL, ACTIVE, ARCHIVED)
}

type ReportUsersFilter struct {
	IDs      []string `json:"ids"`
	Contains string   `json:"contains"` // (CONTAINS, DOES_NOT_CONTAIN)
	Status   string   `json:"status"`   // (ALL, ACTIVE, INACTIVE)
}

type SummaryReportRequest struct {
	DateRangeStart *time.Time         `json:"dateRangeStart"`
	DateRangeEnd   *time.Time         `json:"dateRangeEnd"`
	UsersFilter    *ReportUsersFilter `json:"users"`
	// userGroups
	// clients
	// projects
	// tasks
	TagsFilter         *ReportTagsFilter `json:"tags"`
	IsBillable         *bool             `json:"billable"`
	Description        *string           `json:"description"`
	WithoutDescription *bool             `json:"withoutDescription"`
	SummaryFilter      *SummaryFilter    `json:"summaryFilter"`
	SortOrder          *string           `json:"sortOrder"`  // (ASCENDING, DESCENDING)
	ExportType         *string           `json:"exportType"` // (JSON, CSV, XLSX, PDF)
	Rounding           bool              `json:"rounding"`
	AmountShow         *string           `json:"amountShown"` // (HIDE_AMOUNT, EARNINGS, COST, PROFIT)
	// customFields
}

type SummaryReportTotal struct {
	ID                string   `json:"_id"`
	TotalTime         uint     `json:"totalTime"`
	TotalBillableTime *uint    `json:"totalBillableTime"`
	EntriesCount      uint     `json:"entriesCount"`
	TotalAmount       *float64 `json:"totalAmount"`
}

type SummaryReport struct {
	Chart    []interface{}        `json:"chart"`
	Totals   []SummaryReportTotal `json:"totals"`
	GroupOne []interface{}        `json:"groupOne"`
}

func (r *ReportService) Summary(workspaceID string, opts SummaryReportRequest) (*SummaryReport, error) {
	path := fmt.Sprintf("workspaces/%s/reports/summary", workspaceID)
	req, err := r.client.newReportRequest("POST", path, "", opts)
	if err != nil {
		return nil, err
	}

	var report SummaryReport
	_, err = r.client.do(req, &report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
