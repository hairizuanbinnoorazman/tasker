// Package asana provides the functionality to access the asana API to manipulate and obtain the list from the tool
package asana

import "github.com/hairizuanbinnoorazman/tasker/generic"

var asanaURL = "https://app.asana.com/api/1.0/"

type Asana struct {
	Token string
}

func (a Asana) ListProjects() ([]generic.Project, error) {
	token := a.Token
	err := ListProjects(token)
	return []generic.Project{}, err
}

func (a Asana) ListTasks(projectID string) ([]generic.Task, error) {
	token := a.Token
	err := ListTasks(token, projectID)
	return []generic.Task{}, err
}

func (a Asana) CreateTask(projectID, name string) (generic.Task, error) {
	token := a.Token
	err := CreateTask(token, name, projectID)
	return generic.Task{}, err
}
