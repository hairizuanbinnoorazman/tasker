package asana

import (
	"encoding/json"
	"io/ioutil"
)

func listTags(token, projectID string) ([]genericItem, error) {
	tagsURL := "https://app.asana.com/api/1.0/tags"

	resp, respErr := genericHTTPGet(token, tagsURL)
	if respErr != nil {
		return nil, respErr
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tags genericList
	errJSON := json.Unmarshal(body, &tags)
	if errJSON != nil {
		return nil, errJSON
	}
	return tags.Data, nil
}
