package githubservice

type Repository struct {
	Name        string `json:"name"`
	Language    string `json:"language"`
	Stars       int    `json:"stargazers_count"`
	LastUpdated string `json:"updated_at"`
}

// User represents a GitHub user.
type User struct {
	Followers int `json:"followers"`
}

type Language string

const (
	CSharp     Language = "csharp"
	Go         Language = "go"
	Python     Language = "python"
	Java       Language = "java"
	JavaScript Language = "javascript"
	// Add more languages as needed
)
