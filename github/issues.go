package github

import (
	"fmt"
	"io/ioutil"
)

func listIssues(token, projectID string) error {
	issueURL := githubURL + fmt.Sprintf("/repos/%v/issues", projectID)
	resp, err := genericHTTPGet(token, issueURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return nil
}
