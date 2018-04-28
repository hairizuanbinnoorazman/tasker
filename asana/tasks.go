package asana

type task struct {
	ID             int           `json:"id"`
	Assignee       genericItem   `json:"assignee"`
	AssigneeStatus string        `json:"assignee_status"`
	Completed      bool          `json:"completed"`
	Name           string        `json:"name"`
	Notes          string        `json:"notes"`
	Tags           []genericItem `json:"tags"`
}
