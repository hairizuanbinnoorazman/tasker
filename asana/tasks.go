package asana

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type task struct {
	ID             int           `json:"id"`
	Assignee       genericItem   `json:"assignee"`
	AssigneeStatus string        `json:"assignee_status"`
	Completed      bool          `json:"completed"`
	Name           string        `json:"name"`
	Notes          string        `json:"notes"`
	Tags           []genericItem `json:"tags"`
}

func completeTask(token, taskID string) error {
	taskURL := asanaURL + fmt.Sprintf("tasks/%v", taskID)

	type updateTaskField struct {
		Completed bool `json:"completed"`
	}

	type updateTaskRequest struct {
		Data updateTaskField `json:"data"`
	}

	reqBody := updateTaskRequest{Data: updateTaskField{Completed: true}}
	reqJSON, _ := json.Marshal(reqBody)
	resp, err := genericHTTPPut(token, taskURL, bytes.NewReader(reqJSON))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func updateAssignee(token, projectID string) error {
	return nil
}

func updateLabel(token, projectID string) error {
	return nil
}
