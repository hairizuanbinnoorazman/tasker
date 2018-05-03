// Package github provides the functionality to access the asana API to manipulate and obtain the list from the tool
//
// A point to take note:
// The ID that is return in github responses represent nothing; The reference to issues and repos still take into
// account of github's url structure of organization and the repo's name.
package github

import "github.com/hairizuanbinnoorazman/tasker/generic"

var githubURL = "https://api.github.com"

// Service struct defines the data tokens required to allow the data to be pulled out
// tokens are not tied to the function signatures as there is no guarantee that the
// different task management platforms utilize the same way of authenticating with the same
// kind of values
type Service struct {
	token string
}

// ListProjects in the context of the Github tasker package revolves around the concept of repos
// This command would list the repos that users would be able to access to
func (a Service) ListProjects() ([]generic.Project, error) {
	return []generic.Project{}, nil
}

// ListTasks in the context of the Github tasker package revolves around the concept of issues
// This command would list down list of issues of the repo.
func (a Service) ListTasks(projectID string) ([]generic.Task, error) {
	return []generic.Task{}, nil
}

// CreateTask in the context of the Github tasker package revolves aroudn the concept of creating new issues
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
