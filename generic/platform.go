package generic

// PlatformActions lists all actions that should be provided by a task management integration
// In the case that the value is not meant to be available, we can think of providing defaults that would indicate so
type PlatformActions interface {
	ListProjects() ([]Project, error)
	ListTasks(projectID string) ([]Task, error)
	ListUsers(projectID string) ([]User, error)
	CreateTask(projectID, taskName string) (Task, error)
	CompleteTask(projectID, taskID string) (Task, error)
	AssignTask(projectID, taskID, userID string) (Task, error)
	AssignLabel(projectID, taskID, label string) (Task, error)
}

// ListProjects is a generic function that runs the platform's list project command
func ListProjects(actions PlatformActions) ([]Project, error) {
	return actions.ListProjects()
}

// ListTasks is a generic function that runs the platform's list tasks command
func ListTasks(actions PlatformActions, projectID string) ([]Task, error) {
	return actions.ListTasks(projectID)
}

// ListUsers is a generic function that runs the platform's list users command based of Project
func ListUsers(actions PlatformActions, projectID string) ([]User, error) {
	return actions.ListUsers(projectID)
}

// CreateTask is a generic function that runs the platform's create task command
func CreateTask(actions PlatformActions, projectID, taskName string) (Task, error) {
	return actions.CreateTask(projectID, taskName)
}

// CompleteTask is a generic function that runs the platform's complete task command
func CompleteTask(actions PlatformActions, projectID, taskID string) (Task, error) {
	return actions.CompleteTask(projectID, taskID)
}

// AssignTask is a generic function that runs the platform's assign task command
func AssignTask(actions PlatformActions, projectID, taskID, userID string) (Task, error) {
	return actions.AssignTask(projectID, taskID, userID)
}

// AssignLabel is a generic function that runs the platform's assign label command
func AssignLabel(actions PlatformActions, projectID, taskID, label string) (Task, error) {
	return actions.AssignLabel(projectID, taskID, label)
}
