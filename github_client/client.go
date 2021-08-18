package githubclient

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/trunglv2018/backend-coding-test/common"
	"github.com/trunglv2018/backend-coding-test/helpers"
	"github.com/trunglv2018/backend-coding-test/models"
)

var tagLog = "[github_api]"

type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

func GetOfficialRepositories(userId string) ([]*models.OfficialGithubRepository, error) {
	var repositories []*models.OfficialGithubRepository
	resp, err := http.Get(fmt.Sprintf(common.GetRepositories, userId))
	if err != nil {
		log.Println(tagLog, err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		var errResponse *ErrorResponse
		err = helpers.UnmarshalResponseBody(resp.Body, &errResponse)
		if err != nil {
			return nil, err
		}
		if errResponse != nil {
			return nil, errors.New(errResponse.Message)
		}
		return nil, errors.New(common.ErrUnspecify)
	}

	return repositories, helpers.UnmarshalResponseBody(resp.Body, &repositories)
}

func GetRepositoriesFromOfficial(officialGithubRepositories []*models.OfficialGithubRepository) []*models.Repository {
	var repositories = []*models.Repository{}
	for _, OfficialGithubRepository := range officialGithubRepositories {
		repositories = append(repositories, OfficialGithubRepository.ToRepository())
	}
	return repositories
}
