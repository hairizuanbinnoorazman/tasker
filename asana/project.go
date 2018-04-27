package asana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type project struct {
	Name  string      `json:"name"`
	ID    string      `json:"id"`
	Owner genericItem `json:"owner"`
	Team  genericItem `json:"team"`
}

func getProject(token, projectID string) (project, error) {
	projectURL := asanaURL + fmt.Sprintf("projects/%s", projectID)
	resp, err := genericHTTPGet(token, projectURL)
	if err != nil {
		return project{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	type response struct {
		Data project `json:"data"`
	}

	var parsedResponse response
	errJSON := json.Unmarshal(body, &parsedResponse)
	if errJSON != nil {
		return project{}, errJSON
	}
	return parsedResponse.Data, nil
}
