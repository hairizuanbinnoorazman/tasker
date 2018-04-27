package asana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type genericList struct {
	Data []genericItem `json:"data"`
}

type genericItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func listTasks(token, project string) error {
	projectTaskURL := asanaURL + fmt.Sprintf("projects/%s/tasks", project)
	resp, err := genericHTTPGet(token, projectTaskURL)
	if err != nil {
		fmt.Println("Error in getting response")
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var parsedResponse genericList
	errJSON := json.Unmarshal(body, &parsedResponse)
	if errJSON != nil {
		fmt.Println("Error in unmarshalling content")
		return errJSON
	}
	value, _ := json.MarshalIndent(parsedResponse, "", "\t")
	fmt.Println(string(value))
	return nil
}
