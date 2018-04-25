// Package asana provides the functionality to access the asana API to manipulate and obtain the list from the tool
package trello

import "github.com/hairizuanbinnoorazman/tasker/generic"

type Trello struct{}

func (a Trello) ListProjects() ([]generic.Project, error) {
	return []generic.Project{}, nil
}

func (a Trello) ListTasks(projectID string) ([]generic.Task, error) {
	return []generic.Task{}, nil
}

func (a Trello) CreateTask(projectID, name string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Trello) ListUsers(projectID string) ([]generic.User, error) {
	return []generic.User{}, nil
}

func (a Trello) CompleteTask(projectID, taskID string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Trello) AssignTask(projectID, taskID, userID string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Trello) AssignLabel(projectID, taskID, label string) (generic.Task, error) {
	return generic.Task{}, nil
}
