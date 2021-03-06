package asana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type genericRequest struct {
	Data taskRequest `json:"data"`
}

type taskRequest struct {
	Projects string `json:"projects"`
	Name     string `json:"name"`
}

func createTask(token, name, project string) error {
	projectTaskURL := asanaURL + "tasks"

	client := &http.Client{}

	taskRequestInstance := taskRequest{Projects: project, Name: name}
	reqBody := genericRequest{Data: taskRequestInstance}
	reqBodyByte, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", projectTaskURL, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		fmt.Println("Error creating new request")
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in getting response")
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}
