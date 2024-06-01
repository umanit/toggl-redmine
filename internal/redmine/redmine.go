package redmine

type TimeEntriesList struct {
	TimeEntries []TimeEntry `json:"time_entries"`
}

type TimeEntry struct {
	Issue struct {
		Id int `json:"id"`
	} `json:"issue"`
	Hours   float64 `json:"hours"`
	SpentOn string  `json:"spent_on"`
}
