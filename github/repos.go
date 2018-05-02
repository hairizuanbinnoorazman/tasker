package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type simpleRepo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func listRepos(token string) error {
	repoURL := githubURL + "/user/repos"
	resp, err := genericHTTPGet(token, repoURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var parsedResponse []simpleRepo
	errJSON := json.Unmarshal(body, &parsedResponse)
	if errJSON != nil {
		fmt.Println("Error in unmarshalling content")
		return errJSON
	}
	value, _ := json.MarshalIndent(parsedResponse, "", "\t")
	fmt.Println(string(value))
	return nil
}
