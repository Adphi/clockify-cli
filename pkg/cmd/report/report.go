package report

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/pkuebler/clockify-cli/pkg/clockify"
	"github.com/pkuebler/clockify-cli/pkg/runtime"
)

const (
	layoutISO = "2006-01-02"
)

type ReportConfig struct {
	Range string
}

// NewCmdReport returns new initialized instance of report sub command
func NewCmdReport(ctx *runtime.Runtime) *cobra.Command {
	reportConfig := ReportConfig{}

	cmd := &cobra.Command{
		Use:     "report",
		Short:   "Report worktime",
		Long:    "Create a workreport based on workday settings.",
		Example: "report --range currentyear, currentweek, week, month or year",
		Run: func(cmd *cobra.Command, args []string) {
			// start and end
			now := time.Now()
			currentYear, _, _ := now.Date()
			currentLocation := now.UTC().Location()
			currentMonth := now.UTC().Month()

			var startDate time.Time
			var endDate time.Time

			if reportConfig.Range == "week" {
				startDate = now.UTC()
				endDate = now.UTC()
				startDate = startDate.AddDate(0, 0, -6)
			} else if reportConfig.Range == "month" {
				startDate = now.UTC()
				endDate = now.UTC()
				startDate = startDate.AddDate(0, -1, 0)
			} else if reportConfig.Range == "year" {
				startDate = now.UTC()
				endDate = now.UTC()
				startDate = startDate.AddDate(-1, 0, 0)
			} else if reportConfig.Range == "currentweek" {
				startDate = now.UTC()
				endDate = startDate.UTC()

				if wd := startDate.Weekday(); wd == time.Sunday {
					startDate = startDate.AddDate(0, 0, -6)
				} else {
					startDate = startDate.AddDate(0, 0, -int(wd)+1)
					endDate = endDate.AddDate(0, 0, 7-int(wd))
				}
				endDate = endDate.AddDate(0, 0, 1).Add(-time.Second)
			} else {
				// current year
				startDate = time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentLocation)
				endDate = time.Date(currentYear, currentMonth, now.Day()+1, 0, 0, 0, 0, currentLocation).Add(-time.Second)
			}

			// set workdays
			calendar := NewCalendar()
			calendar.PlanWorkhours(time.Monday, ctx.Config.ReportConfig.Workhours.Monday)
			calendar.PlanWorkhours(time.Tuesday, ctx.Config.ReportConfig.Workhours.Tuesday)
			calendar.PlanWorkhours(time.Wednesday, ctx.Config.ReportConfig.Workhours.Wednesday)
			calendar.PlanWorkhours(time.Thursday, ctx.Config.ReportConfig.Workhours.Thursday)
			calendar.PlanWorkhours(time.Friday, ctx.Config.ReportConfig.Workhours.Friday)
			calendar.PlanWorkhours(time.Saturday, ctx.Config.ReportConfig.Workhours.Saturday)
			calendar.PlanWorkhours(time.Sunday, ctx.Config.ReportConfig.Workhours.Sunday)

			// request user data
			user, err := ctx.Client.User.Info()
			if err != nil {
				log.Fatal(err)
				return
			}

			// workspace
			if ctx.WorkspaceID == "" {
				ctx.WorkspaceID = user.DefaultWorkspace
			}

			pageSize := 50
			finish := false
			end := endDate
			lastID := ""

			ctx.Log.Debugf("Start at %v...", end)

			for !finish {
				opts := &clockify.TimeEntryListFilter{
					Start:    &startDate,
					End:      &end,
					PageSize: pageSize,
				}

				ctx.Log.Tracef("Request next page at %v - %v...", startDate, end)

				entries, err := ctx.Client.TimeEntry.List(ctx.WorkspaceID, user.ID, opts)
				if err != nil {
					log.Fatal(err)
				}

				ctx.Log.Tracef("Get %d entries", len(*entries))

				newEntries := false
				for _, entry := range *entries {
					if lastID != "" && !newEntries {
						if entry.ID != lastID {
							continue
						} else {
							newEntries = true
							continue
						}
					}

					if entry.UserID != user.ID {
						ctx.Log.Trace("not from current user!")
						continue
					}

					entryEnd := entry.TimeInterval.End
					if entryEnd == nil {
						entryEnd = &now
					}

					day := entry.TimeInterval.Start.YearDay()
					duration := entryEnd.Sub(entry.TimeInterval.Start)
					calendar.Work(day, duration)

					ctx.Log.Tracef("%s - %s (%s/%s)", entry.TimeInterval.Start, entry.Description, calendar.WorkedHours(entry.TimeInterval.Start).String(), calendar.Workhours(entry.TimeInterval.Start).String())
				}

				finish = len(*entries) < pageSize
				end = (*entries)[len(*entries)-10].TimeInterval.Start
				lastID = (*entries)[len(*entries)-1].ID

				currentState := calendar.Report(startDate, endDate)
				ctx.Log.Tracef("Request Resume: %s / %s", currentState.Sum.WorkedHoursReadable, currentState.Sum.WorkhoursReadable)
			}

			output, _ := json.MarshalIndent(calendar.Report(startDate, endDate), "", "\t")
			ctx.Log.Debug(string(output))
			fmt.Println(string(output))
		},
	}

	cmd.Flags().StringVarP(&reportConfig.Range, "range", "r", "currentyear", "--range currentyear")

	return cmd
}
