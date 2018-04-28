// Package asana provides the functionality to access the asana API to manipulate and obtain the list from the tool
package asana

import "github.com/hairizuanbinnoorazman/tasker/generic"

var asanaURL = "https://app.asana.com/api/1.0/"

// Asana struct defines the data tokens required to allow the data to be pulled out
// tokens are not tied to the function signatures as there is no guarantee that the
// different task management platforms utilize the same way of authenticating with the same
// kind of values
type Asana struct {
	Token string
}

// ListProjects implements an internal call to list all projects in Asana based on an account
func (a Asana) ListProjects() ([]generic.Project, error) {
	token := a.Token
	err := listProjects(token)
	return []generic.Project{}, err
}

// ListTasks implements an internal call to list all tasks in Asana based on an account
func (a Asana) ListTasks(projectID string) ([]generic.Task, error) {
	token := a.Token
	err := listTasks(token, projectID)
	return []generic.Task{}, err
}

// CreateTask implements an internal call to create task in Asana
func (a Asana) CreateTask(projectID, name string) (generic.Task, error) {
	token := a.Token
	err := createTask(token, name, projectID)
	return generic.Task{}, err
}

// ListUsers implements an internal call to list users in Asana for a project
func (a Asana) ListUsers(projectID string) ([]generic.User, error) {
	token := a.Token
	err := listUsersByProject(token, projectID)
	return []generic.User{}, err
}

// CompleteTask implements an internal call to set a task as complete in Asana
func (a Asana) CompleteTask(projectID, taskID string) (generic.Task, error) {
	token := a.Token
	completeTask(token, taskID)
	return generic.Task{}, nil
}

// AssignTask implements an internal call to assign tasks to a user
// Requires IDs of project, task, user
func (a Asana) AssignTask(projectID, taskID, userID string) (generic.Task, error) {
	return generic.Task{}, nil
}

// AssignLabel implements an internal call to assign labels to task
func (a Asana) AssignLabel(projectID, taskID, label string) (generic.Task, error) {
	return generic.Task{}, nil
}
