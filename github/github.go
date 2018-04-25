// Package github provides the functionality to access the asana API to manipulate and obtain the list from the tool
package github

import "github.com/hairizuanbinnoorazman/tasker/generic"

type Github struct{}

func (a Github) ListProjects() ([]generic.Project, error) {
	return []generic.Project{}, nil
}

func (a Github) ListTasks(projectID string) ([]generic.Task, error) {
	return []generic.Task{}, nil
}

func (a Github) CreateTask(projectID, name string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Github) ListUsers(projectID string) ([]generic.User, error) {
	return []generic.User{}, nil
}

func (a Github) CompleteTask(projectID, taskID string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Github) AssignTask(projectID, taskID, userID string) (generic.Task, error) {
	return generic.Task{}, nil
}

func (a Github) AssignLabel(projectID, taskID, label string) (generic.Task, error) {
	return generic.Task{}, nil
}
