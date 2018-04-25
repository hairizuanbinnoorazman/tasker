package generic

// PlatformActions lists all actions that should be provided by a task management integration
// In the case that the value is not meant to be available, we can think of providing defaults that would indicate so
type PlatformActions interface {
	ListProjects() ([]Project, error)
	ListTasks(projectID string) ([]Task, error)
	ListUsers(projectID string) ([]User, error)
	CreateTask(projectID, taskName string) (Task, error)
	// CompleteTask(projectID, taskID string) (Task, error)
	// AssignTask(projectID, taskID, userID string) (Task, error)
	// AssignLabel(projectID, taskID, label string) (Task, error)
}

func ListProjects(actions PlatformActions) ([]Project, error) {
	return actions.ListProjects()
}

func ListTasks(actions PlatformActions, projectID string) ([]Task, error) {
	return actions.ListTasks(projectID)
}

func ListUsers(actions PlatformActions, projectID string) ([]User, error) {
	return actions.ListUsers(projectID)
}

func CreateTask(actions PlatformActions, projectID, taskName string) (Task, error) {
	return actions.CreateTask(projectID, taskName)
}
