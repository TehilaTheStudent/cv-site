// main.go: The entry point for the application.
package main

import (
	"log"

	"github.com/TehilaTheStudent/cv-site/githubservice"
	"github.com/TehilaTheStudent/cv-site/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	// Load environment variables from .env file in development
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	// Set default environment
	viper.AutomaticEnv()
	viper.SetDefault("GITHUB_TOKEN", "")

	token := viper.GetString("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}
	client := githubservice.NewGitHubClient(&token)
	r := gin.Default()
	appInstance := &routes.App{GitHubClient: client}
	routes.PortfolioRoute(r, appInstance)
	routes.SearchRoute(r, appInstance)
	r.Run(":8080")
}
