package api

import (
	"github.com/gin-gonic/gin"
	githubclient "github.com/trunglv2018/backend-coding-test/github_client"
	"github.com/trunglv2018/backend-coding-test/models"
)

type RepositoryApi struct {
	*gin.RouterGroup
}

func NewRepositoryApi(route *gin.RouterGroup) {
	var repositoryApi = &RepositoryApi{
		RouterGroup: route,
	}
	repositoryApi.GET(":userId/repositories", repositoryApi.GetRepositories)
}

func (s *RepositoryApi) GetRepositories(c *gin.Context) {
	var userId = c.Param("userId")
	var officialGithubRepositories, err = githubclient.GetOfficialRepositories(userId)
	if err != nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	} else {
		var repositories []*models.Repository
		if officialGithubRepositories != nil {
			repositories = githubclient.GetRepositoriesFromOfficial(officialGithubRepositories)
		}
		c.JSON(200, repositories)
	}
}
