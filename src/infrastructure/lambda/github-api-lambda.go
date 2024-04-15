package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
)

type GitHubRequest struct {
	Repository string `json:"repository"`
}

type GitHubResponse struct {
	Repository  string `json:"repository"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
}

func GithubHandler(ctx context.Context, request GitHubRequest) (*GitHubResponse, error) {
	response, err := fetchGitHubData(request.Repository)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func fetchGitHubData(repository string) (*GitHubResponse, error) {
	githubBaseUrl := os.Getenv("GITHUB_API_BASE_URL")
	url := fmt.Sprintf(githubBaseUrl+"/%s", repository)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching GitHub data: %v", err)
	}
	defer resp.Body.Close()

	var repoData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&repoData)
	if err != nil {
		return nil, fmt.Errorf("error decoding GitHub data: %v", err)
	}

	response := &GitHubResponse{
		Repository:  repository,
		Description: repoData["description"].(string),
		Stars:       int(repoData["stargazers_count"].(float64)),
	}

	return response, nil
}

func main() {
	lambda.Start(GithubHandler)
}
