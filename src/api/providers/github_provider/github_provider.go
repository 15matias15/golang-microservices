package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/15matias15/golang-microservices/src/api/clients/restclient"
	github2 "github.com/15matias15/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func GetAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github2.CreateRepoRequest) (*github2.CreateRepoResponse, *github2.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, GetAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create a new repo in github %s", err.Error()))
		return nil, &github2.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github2.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github2.GithubErrorResponse

		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github2.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github2.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response %s", err.Error()))
		return nil, &github2.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error unmarshalling github create repo",
		}
	}

	return &result, nil
}

//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
