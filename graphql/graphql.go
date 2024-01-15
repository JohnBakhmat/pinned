package graphql

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Project struct {
	Name        string
	Description string
	Url         string
	Stars       int
	Forks       int
	Language    []string
}

func GetProjects(username string) ([]Project, error) {
	key := os.Getenv("GITHUB_TOKEN")
	if key == "" {
		err := fmt.Errorf("must set GITHUB_TOKEN=<github token>")
		return nil, err
	}

	ctx := context.Background()

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: key},
	)
	httpClient := oauth2.NewClient(ctx, src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		User struct {
			PinnedItems struct {
				Nodes []struct {
					Repository struct {
						Name           string
						Description    string
						URL            string
						ForkCount      int
						StargazerCount int
					} `graphql:"... on Repository"`
				}
			} `graphql:"pinnedItems(first: 6, types: [REPOSITORY])"`
		} `graphql:"user(login: $username)"`
	}

	vars := map[string]interface{}{
		"username": githubv4.String(username),
	}

	err := client.Query(ctx, &query, vars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}

	projects := []Project{}

	for _, repo := range query.User.PinnedItems.Nodes {
		project := Project{
			Name:        repo.Repository.Name,
			Description: repo.Repository.Description,
			Url:         repo.Repository.URL,
			Forks:       repo.Repository.ForkCount,
			Stars:       repo.Repository.StargazerCount,
		}
		projects = append(projects, project)
	}

	return projects, nil
}
