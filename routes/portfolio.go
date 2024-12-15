package routes

import (
	"github.com/gin-gonic/gin"
)

func PortfolioRoute(r *gin.Engine, app *App) {
	r.GET("/portfolio", func(c *gin.Context) {

		username := "TehilaTheStudent"
		repos, err := app.GitHubClient.GetUserRepositories(username)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, repos)
	})
}
