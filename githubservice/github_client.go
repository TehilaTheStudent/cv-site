package githubservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GitHubClient holds the token and provides methods to interact with the GitHub API.
type GitHubClient struct {
	Token   *string
	BaseURL string
}

// NewGitHubClient creates a new GitHub client with the provided token.
func NewGitHubClient(token *string) *GitHubClient {
	return &GitHubClient{
		Token:   token,
		BaseURL: "https://api.github.com",
	}
}

// makeRequest is a helper to make HTTP requests to the GitHub API.
func (client *GitHubClient) makeRequest(method, endpoint string, body interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", client.BaseURL, endpoint)

	var requestBody []byte
	if body != nil && method != "POST" && method != "PUT" {
		return nil, fmt.Errorf("body is not allowed for method %s", method)
	}
	if body != nil {
		var err error
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	if client.Token != nil {
		req.Header.Set("Authorization", "token "+*client.Token)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		remaining := resp.Header.Get("X-RateLimit-Remaining")
		reset := resp.Header.Get("X-RateLimit-Reset")
		return nil, fmt.Errorf("rate limit exceeded. Remaining: %s, Reset at: %s", remaining, reset)
	}
	if resp.StatusCode >= 400 {
		responseBody, _ := io.ReadAll(resp.Body) // Read without handling the error to ensure we log it
		return nil, fmt.Errorf("error: %s, response: %s", resp.Status, string(responseBody))
	}

	return io.ReadAll(resp.Body)
}

// makeRequestNoAuth is a helper to make HTTP requests to the GitHub API without authentication.
func (client *GitHubClient) makeRequestNoAuth(method, endpoint string, body interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", client.BaseURL, endpoint)

	var requestBody []byte
	if body != nil {
		var err error
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
