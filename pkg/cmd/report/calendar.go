package report

import (
	"fmt"
	"time"
)

type Calendar struct {
	plannedWorkingTime map[time.Weekday]time.Duration
	workedHours        map[int]time.Duration            // yearday
	filter             map[string]map[int]time.Duration // filtername, yearday
}

type Report struct {
	Workhours   time.Duration `json:"workhours"`
	WorkedHours time.Duration `json:"worked_hours"`

	WorkhoursReadable   string `json:"workhours_readable"`
	WorkedHoursReadable string `json:"worked_hours_readable"`

	Filter         map[string]time.Duration `json:"filter"`
	FilterReadable map[string]string        `json:"filter_readable"`

	Diff         time.Duration `json:"diff"`
	DiffReadable string        `json:"diff_readable"`
}

func NewReport() *Report {
	return &Report{
		Workhours:   time.Duration(0),
		WorkedHours: time.Duration(0),

		Filter:         map[string]time.Duration{},
		FilterReadable: map[string]string{},
	}
}

func (r *Report) Readable() {
	r.WorkhoursReadable = r.Workhours.String()
	r.WorkedHoursReadable = r.WorkedHours.String()

	r.Diff = r.WorkedHours - r.Workhours
	r.DiffReadable = r.Diff.String()

	for filterName, filter := range r.Filter {
		r.FilterReadable[filterName] = filter.String()
	}
}

type CalendarReport struct {
	Months map[time.Month]*Report `json:"months"`
	Sum    *Report                `json:"sum"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func NewCalendarReport(start time.Time, end time.Time) *CalendarReport {
	return &CalendarReport{
		Sum: NewReport(),

		Months: map[time.Month]*Report{},

		Start: start,
		End:   end,
	}
}

func (c *CalendarReport) AddTime(i time.Time, workhours time.Duration, workedHours time.Duration) {
	if _, ok := c.Months[i.Month()]; !ok {
		c.Months[i.Month()] = NewReport()
	}

	c.Sum.Workhours += workhours
	c.Sum.WorkedHours += workedHours

	c.Months[i.Month()].Workhours += workhours
	c.Months[i.Month()].WorkedHours += workedHours
}

func (c *CalendarReport) AddTimeToFilter(i time.Time, filterName string, workedHours time.Duration) {
	if _, ok := c.Months[i.Month()]; !ok {
		c.Months[i.Month()] = NewReport()
	}

	if _, ok := c.Sum.Filter[filterName]; !ok {
		c.Sum.Filter[filterName] = time.Duration(0)
		c.Months[i.Month()].Filter[filterName] = time.Duration(0)
	}

	c.Sum.Filter[filterName] += workedHours
	c.Months[i.Month()].Filter[filterName] += workedHours
}

func NewCalendar() *Calendar {
	hours, _ := time.ParseDuration("0h")

	return &Calendar{
		plannedWorkingTime: map[time.Weekday]time.Duration{
			time.Monday:    hours,
			time.Tuesday:   hours,
			time.Wednesday: hours,
			time.Thursday:  hours,
			time.Friday:    hours,
			time.Saturday:  hours,
			time.Sunday:    hours,
		},
		workedHours: map[int]time.Duration{},
		filter:      map[string]map[int]time.Duration{},
	}
}

func (c *Calendar) PlanWorkhours(day time.Weekday, hours int) {
	duration, _ := time.ParseDuration(fmt.Sprintf("%dh", hours))

	c.plannedWorkingTime[day] = duration
}

func (c *Calendar) Work(yearDay int, d time.Duration) {
	if _, ok := c.workedHours[yearDay]; !ok {
		c.workedHours[yearDay] = 0
	}

	c.workedHours[yearDay] += d
}

func (c *Calendar) Workhours(date time.Time) time.Duration {
	day := date.Weekday()

	return c.plannedWorkingTime[day]
}

func (c *Calendar) WorkedHours(t time.Time) time.Duration {
	day := t.YearDay()

	if _, ok := c.workedHours[day]; !ok {
		return time.Duration(0)
	}

	return c.workedHours[day]
}

func (c *Calendar) Report(start, end time.Time) *CalendarReport {
	r := NewCalendarReport(start, end)

	var i time.Time
	for i = start; i.Before(end); i = i.AddDate(0, 0, 1) {
		r.AddTime(i, c.Workhours(i), c.WorkedHours(i))
	}

	if i.Equal(end) {
		r.AddTime(i, c.Workhours(i), c.WorkedHours(i))
	}

	r.Sum.Readable()

	for _, month := range r.Months {
		month.Readable()
	}

	return r
}
