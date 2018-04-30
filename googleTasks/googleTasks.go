// Package asana provides the functionality to access the asana API to manipulate and obtain the list from the tool
package googleTasks

import "github.com/hairizuanbinnoorazman/tasker/generic"

type Service struct{}

func (a Service) ListProjects() ([]generic.Project, error) {
	return []generic.Project{}, nil
}

func (a Service) ListTasks(projectID string) ([]generic.Task, error) {
	return []generic.Task{}, nil
}

func (a Service) CreateTask(projectID, name string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Service) ListUsers(projectID string) ([]generic.User, error) {
	return []generic.User{}, nil
}

func (a Service) CompleteTask(projectID, taskID string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Service) AssignTask(projectID, taskID, userID string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Service) AssignLabel(projectID, taskID, label string) (generic.Task, error) {
	return generic.Task{}, nil
}
