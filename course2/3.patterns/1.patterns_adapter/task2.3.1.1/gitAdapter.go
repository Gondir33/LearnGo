package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v53/github"
)

/*
	func main() {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: "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"},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)
		g := NewGithub(client)
		fmt.Println(g.GetGists(context.Background(), "ptflp"))
		fmt.Println(g.GetRepos(context.Background(), "ptflp"))
	}
*/
type Item struct {
	Title       string
	Description string
	Link        string
}

type GithuberAdapter interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type GithubAdapter struct {
	client *github.Client
	Reposer
	Gister
}

//go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=Reposer --inpackage
type Reposer interface {
	List(ctx context.Context, user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}
type Repos struct {
	g *GithubAdapter
}

// go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=Gister --inpackage
type Gister interface {
	List(ctx context.Context, user string, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Gist struct {
	g *GithubAdapter
}

func NewGithub(client *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		client: client,
	}
	g.Gister = &Gist{g}
	g.Reposer = &Repos{g}
	return g
}

func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	items := make([]Item, 0)
	gists, resp, err := g.Gister.List(ctx, username, nil)
	if err != nil {
		return []Item{}, err
	}
	if resp.StatusCode == http.StatusOK {
		for _, gist := range gists {
			items = append(items, Item{
				Title:       gist.GetID(),
				Description: gist.GetDescription(),
				Link:        gist.GetHTMLURL(),
			})
		}
		return items, nil
	}
	return []Item{}, errors.New(fmt.Sprintf("response StatusCode %v", resp.StatusCode))

}

func (g *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	items := make([]Item, 0)
	repos, resp, err := g.Reposer.List(ctx, username, nil)
	if err != nil {
		return []Item{}, err
	}
	if resp.StatusCode == http.StatusOK {
		for _, rep := range repos {
			items = append(items, Item{
				Title:       rep.GetName(),
				Description: rep.GetDescription(),
				Link:        rep.GetHTMLURL(),
			})
		}
		return items, nil
	}
	return []Item{}, errors.New(fmt.Sprintf("response.StatusCode %v", resp.StatusCode))
}

func (r *Repos) List(ctx context.Context, user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	return r.g.client.Repositories.List(ctx, user, opts)
}

func (gist *Gist) List(ctx context.Context, user string, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	return gist.g.client.Gists.List(ctx, user, opts)
}
