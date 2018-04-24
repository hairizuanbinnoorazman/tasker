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

// PlatformActions lists all actions that should be provided by a task management integration
// In the case that the value is not meant to be available, we can think of providing defaults that would indicate so
type PlatformActions interface {
	ListProjects() ([]Project, error)
	ListTasks(projectID string) ([]Task, error)
	ListUsers(projectID string) ([]User, error)
	CreateTask(projectID string) (Task, error)
	CompleteTask(projectID, taskID string) (Task, error)
	AssignTask(projectID, taskID, userID string) (Task, error)
	AssignLabel(projectID, taskID, label string) (Task, error)
}
