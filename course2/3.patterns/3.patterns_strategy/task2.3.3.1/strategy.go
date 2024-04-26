package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v53/github"
)

type Item struct {
	Title       string
	Description string
	Link        string
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GithubGist struct {
	client *github.Client
	Gister
}

func NewGithubGist(client *github.Client) *GithubGist {
	g := &GithubGist{client: client, Gister: &Gist{client: client}}
	return g
}

func (g *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	items := make([]Item, 0)
	gists, response, err := g.Gister.List(ctx, username, nil)
	if err != nil {
		return []Item{}, err
	}
	if response.StatusCode != http.StatusOK {
		return []Item{}, errors.New(fmt.Sprintf("Response code %v", response.StatusCode))
	}
	for _, gist := range gists {
		items = append(items, Item{
			Title:       gist.GetID(),
			Description: gist.GetDescription(),
			Link:        gist.GetHTMLURL(),
		})
	}
	return items, nil
}

type GithubRepo struct {
	client *github.Client
	Reposer
}

func NewGithubRepo(client *github.Client) *GithubRepo {
	r := &GithubRepo{client: client, Reposer: &Repos{client: client}}
	return r
}

func (r *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	items := make([]Item, 0)
	repos, response, err := r.Reposer.List(ctx, username, nil)
	if err != nil {
		return []Item{}, err
	}
	if response.StatusCode != http.StatusOK {
		return []Item{}, errors.New(fmt.Sprintf("Response code %v", response.StatusCode))
	}
	for _, repo := range repos {
		items = append(items, Item{
			Title:       repo.GetName(),
			Description: repo.GetDescription(),
			Link:        repo.GetHTMLURL(),
		})
	}
	return items, nil
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type GeneralGithub struct {
	client *github.Client
}

func NewGeneralGithub(client *github.Client) *GeneralGithub {
	gg := &GeneralGithub{client: client}
	return gg
}

func (gg *GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

// go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=Reposer --inpackage
type Reposer interface {
	List(ctx context.Context, user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type Repos struct {
	client *github.Client
}

func (r *Repos) List(ctx context.Context, user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	return r.client.Repositories.List(ctx, user, opts)
}

// go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=Gister --inpackage
type Gister interface {
	List(ctx context.Context, user string, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Gist struct {
	client *github.Client
}

func (g *Gist) List(ctx context.Context, user string, opts *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	return g.client.Gists.List(ctx, user, opts)
}

/*
func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)

	gg := NewGeneralGithub(client)

	data, err := gg.GetItems(context.Background(), "ptflp", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

	data, err = gg.GetItems(context.Background(), "ptflp", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
*/
