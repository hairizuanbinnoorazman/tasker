package asana

import (
	"encoding/json"
	"errors"
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

func getTagID(token, projectID, tagName string) (int, error) {
	tags, err := listTags(token, projectID)
	if err != nil {
		return 0, err
	}
	for _, val := range tags {
		if tagName == val.Name {
			return val.ID, nil
		}
	}
	return 0, errors.New("Tag not found")
}
