package github

func listRepos(token string) error {
	repoURL := githubURL + "/user/repos"
	resp, err := genericHTTPGet(token, repoURL)
	if err != nil {
		return err
	}
	return nil
}
