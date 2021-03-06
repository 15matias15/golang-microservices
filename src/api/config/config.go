package config

import "os"

const (
	apiGitHubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGitHubAccessToken)
)

func GetGitHubAccessToken() string {
	return githubAccessToken
}
