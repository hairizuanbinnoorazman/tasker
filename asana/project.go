package asana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getProject(token, project string) error {
	projectURL := asanaURL + fmt.Sprintf("projects/%s", project)

	client := &http.Client{}

	req, err := http.NewRequest("GET", projectURL, nil)
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

	type projectDetails struct {
		Team genericItem `json:"team"`
	}

	type overallProjectDetails struct {
		Data projectDetails `json:"data"`
	}

	var parsedResponse overallProjectDetails
	errJSON := json.Unmarshal(body, &parsedResponse)
	if errJSON != nil {
		fmt.Println("Error in unmarshalling content")
		return errJSON
	}
	value, _ := json.MarshalIndent(parsedResponse, "", "\t")
	fmt.Println(string(value))
	return nil
}
