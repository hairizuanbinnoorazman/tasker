package github

import "net/http"

func genericHTTPGet(token, url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
