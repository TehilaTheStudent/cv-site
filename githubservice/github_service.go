package githubservice

import (
	"encoding/json"
	"fmt"
)

// GetUser fetches a GitHub user by username.
func (client *GitHubClient) GetUser(username string) ([]byte, error) {
	endpoint := fmt.Sprintf("/users/%s", username)
	return client.makeRequestNoAuth("GET", endpoint, nil)
}

// GetUserFollowers fetches the follower count for a specific user.
func (client *GitHubClient) GetUserFollowers(userName string) (int, error) {
	data, err := client.makeRequestNoAuth("GET", "/users/"+userName, nil)
	if err != nil {
		return 0, err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return 0, err
	}

	return user.Followers, nil
}

// SearchRepositoriesByLanguage searches repositories written in the specified language.
func (client *GitHubClient) SearchRepositoriesByLanguage(language Language) ([]Repository, error) {
	query := fmt.Sprintf(`?q=language:%s`, language)
	data, err := client.makeRequest("GET", "/search/repositories"+query,nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Items []Repository `json:"items"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result.Items, nil
}

// GetUserRepositories fetches repositories for a specific user.
func (client *GitHubClient) GetUserRepositories(userName string) ([]Repository, error) {
	data, err := client.makeRequest("GET", fmt.Sprintf("/users/%s/repos", userName), nil)
	if err != nil {
		return nil, err
	}

	var repos []Repository
	if err := json.Unmarshal(data, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}
