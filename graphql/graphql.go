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
	Languages   []string
}

func GetProjects(username string) ([]Project, error) {
	key := os.Getenv("GITHUB_TOKEN")
	fmt.Println(key);
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
						Languages      struct {
							Nodes []struct {
								Name string
							}
						} `graphql:"languages(first:5)"`
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
		languages := make([]string, len(query.User.PinnedItems.Nodes))
		for i, l := range repo.Repository.Languages.Nodes {
			languages[i] = l.Name
		}
		filterEmptyStrings(&languages)

		project := Project{
			Name:        repo.Repository.Name,
			Description: repo.Repository.Description,
			Url:         repo.Repository.URL,
			Forks:       repo.Repository.ForkCount,
			Stars:       repo.Repository.StargazerCount,
			Languages:   languages,
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func filterEmptyStrings(s *[]string) {
	i := 0
	for _, str := range *s {
		if str != "" {
			(*s)[i] = str
			i++
		}
	}
	*s = (*s)[:i]
}
