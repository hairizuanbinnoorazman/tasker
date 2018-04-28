package asana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func listUsersByTeam(token, teamID string) error {
	teamURL := asanaURL + fmt.Sprintf("teams/%v/users", teamID)
	resp, errResp := genericHTTPGet(token, teamURL)
	defer resp.Body.Close()
	if errResp != nil {
		return errResp
	}
	body, _ := ioutil.ReadAll(resp.Body)

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

// listUsers function in Asana requires a 2 step process to retrive users
// Part 1: Obtain the team ID
// Part 2: From the team ID, list all members within that group
func listUsersByProject(token, projectID string) error {
	project, err := getProject(token, projectID)
	if err != nil {
		return err
	}
	teamID := project.Team.ID
	listUsersByTeam(token, strconv.Itoa(teamID))
	return nil
}
