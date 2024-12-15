package routes

import (
	"net/http"

	"github.com/TehilaTheStudent/cv-site/githubservice"
	"github.com/gin-gonic/gin"
)

func SearchRoute(r *gin.Engine, app *App) {
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
			return
		}

		results, err := app.GitHubClient.SearchRepositoriesByLanguage(githubservice.Language(query))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, results)
	})

	r.GET("/followers", func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'username' is required"})
			return
		}

		followers, err := app.GitHubClient.GetUserFollowers(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, followers)
	})
}
