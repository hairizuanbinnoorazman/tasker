package generic

import (
	"time"
)

// Task struct defines a generic task interface that could possibly be used across the various platforms
// Some things to note here:
// - ID is not set to integer in the case where IDs may contain a string instead in the task management platform (e.g. UUID)
// - Label may actually be an identified properly with ID but it may be unnecessary exposure of the platform to the user
//   Instead we can enforce it such that by default a new label will not be created; but user can force its creation
// - Status may actually be an identified properly with ID but it may be unnecessary exposure of the platform to the user
//   Instead we can enforce it such that by default a new label will not be created; but user can force its creation
type Task struct {
	ID          string
	Title       string
	Description string
	StartDate   time.Time
	Deadline    time.Time
	Assigned    string
	Label       string
	Status      string
	Subscribers []User
}
