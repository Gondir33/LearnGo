package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestGithub_GetRepos(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	val := "1"
	desk := "2"
	url := "3"

	repMock := NewMockReposer(t)
	gt := github.NewClient(&http.Client{})
	reposer := NewGithubRepo(gt)
	reposer.Reposer = repMock
	g := NewGeneralGithub(gt)

	tests := []struct {
		name    string
		g       *GeneralGithub
		args    args
		want    []Item
		wantErr bool
		list    []*github.Repository
		resp    *github.Response
		errList error
	}{{
		name:    "without error",
		g:       g,
		args:    args{ctx: context.Background(), username: "Gondir33"},
		want:    []Item{{Title: "1", Description: "2", Link: "3"}},
		wantErr: false,
		list:    []*github.Repository{{Name: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 200}},
		errList: nil,
	}, {
		name:    "with error",
		g:       g,
		args:    args{ctx: context.Background(), username: ""},
		want:    []Item{},
		wantErr: true,
		list:    []*github.Repository{{Name: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 100}},
		errList: nil,
	}, {
		name:    "with error",
		g:       g,
		args:    args{username: ""},
		want:    []Item{},
		wantErr: true,
		list:    []*github.Repository{{Name: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 200}},
		errList: errors.New("ctx can not be nil"),
	}}

	for _, tt := range tests {
		repMock.On("List", tt.args.ctx, tt.args.username, mock.AnythingOfType("*github.RepositoryListOptions")).Return(tt.list, tt.resp, tt.errList)
		got, err := tt.g.GetItems(tt.args.ctx, tt.args.username, reposer)
		if (err != nil) != tt.wantErr {
			t.Errorf("GeneralGithub.GetItems error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Equal(t, tt.want, got)
	}
}

func TestGithub_GetGists(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	val := "1"
	desk := "2"
	url := "3"

	gisMock := NewMockGister(t)
	gt := github.NewClient(&http.Client{})
	gister := NewGithubGist(gt)
	gister.Gister = gisMock
	g := NewGeneralGithub(gt)

	tests := []struct {
		name    string
		g       *GeneralGithub
		args    args
		want    []Item
		wantErr bool
		list    []*github.Gist
		resp    *github.Response
		errList error
	}{{
		name:    "mock",
		g:       g,
		args:    args{ctx: context.Background(), username: ",fapieojewko[dksp]"},
		want:    []Item{{Title: "1", Description: "2", Link: "3"}},
		wantErr: false,
		list:    []*github.Gist{{ID: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 200}},
		errList: nil,
	}, {
		name:    "with error",
		g:       g,
		args:    args{ctx: context.Background(), username: "asd"},
		want:    []Item{},
		wantErr: true,
		list:    []*github.Gist{{ID: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 401}},
		errList: nil,
	}, {
		name:    "with error",
		g:       g,
		args:    args{username: "asd"},
		want:    []Item{},
		wantErr: true,
		list:    []*github.Gist{{ID: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 200}},
		errList: errors.New("ctx can not be nil"),
	}}

	for _, tt := range tests {
		gisMock.On("List", tt.args.ctx, tt.args.username, mock.AnythingOfType("*github.GistListOptions")).Return(tt.list, tt.resp, tt.errList)
		got, err := tt.g.GetItems(tt.args.ctx, tt.args.username, gister)
		if (err != nil) != tt.wantErr {
			t.Errorf("GithubAdapter.GetGists() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Equal(t, tt.want, got)
	}
}

func TestList(t *testing.T) {
	client := github.NewClient(&http.Client{})
	g := &Gist{client: client}
	_, _, err := g.List(nil, "sad", nil)
	if err == nil {
		t.Errorf("should be error")
	}
	r := &Repos{client: client}
	_, _, err = r.List(nil, "sad", nil)
	if err == nil {
		t.Errorf("should be error")
	}
}
