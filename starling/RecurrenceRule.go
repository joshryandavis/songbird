package starling

type RecurrenceRule struct {
	StartDate string   `json:"startDate"`
	Frequency string   `json:"frequency"`
	Interval  int      `json:"interval"`
	Count     int      `json:"count"`
	UntilDate string   `json:"untilDate"`
	WeekStart string   `json:"weekStart"`
	Days      []string `json:"days"`
	MonthDay  int      `json:"monthDay"`
	MonthWeek int      `json:"monthWeek"`
}
