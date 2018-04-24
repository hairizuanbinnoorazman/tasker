package generic

// Platform is a generic struct to describe a platform being used to manage tasks
type Platform struct {
	Name string
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

func (p Platform) ListProjects(actions PlatformActions) ([]Project, error) {
	return actions.ListProjects()
}
