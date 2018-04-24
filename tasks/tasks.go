package tasks

import (
	"time"
)

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

type PlatformActions interface {
	ListProjects() ([]Project, error)
	ListTasks(projectID string) ([]Task, error)
	ListUsers(projectID string) ([]User, error)
	CreateTask(projectID string) (Task, error)
	CompleteTask(projectID, taskID string) (Task, error)
	AssignTask(projectID, taskID, userID string) (Task, error)
	AssignLabel(projectID, taskID, label string) (Task, error)
}
