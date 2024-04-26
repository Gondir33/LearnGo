package main

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestNewGithub(t *testing.T) {
	client := &github.Client{}
	want := NewGithub(client)
	if got := NewGithub(client); !reflect.DeepEqual(got, want) {
		t.Errorf("NewGithub() = %v, want %v", got, want)
	}
}

func TestGithub_GetRepos(t *testing.T) {
	repMock := NewMockReposer(t)
	gt := github.NewClient(&http.Client{})
	g := NewGithub(gt)
	g.Reposer = repMock
	type args struct {
		ctx      context.Context
		username string
	}
	val := "1"
	desk := "2"
	url := "3"
	tests := []struct {
		name    string
		g       *GithubAdapter
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
		name:    "with cache",
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
		if tt.name != "with cache" {
			repMock.On("List", tt.args.ctx, tt.args.username, mock.AnythingOfType("*github.RepositoryListOptions")).Return(tt.list, tt.resp, tt.errList)
		}
		got, err := tt.g.GetRepos(tt.args.ctx, tt.args.username)
		if (err != nil) != tt.wantErr {
			t.Errorf("GithubAdapter.GetRepos() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Equal(t, tt.want, got)
	}
}

func TestGithub_GetGists(t *testing.T) {
	gisMock := NewMockGister(t)
	gt := github.NewClient(&http.Client{})
	g := NewGithub(gt)
	g.Gister = gisMock
	type args struct {
		ctx      context.Context
		username string
	}
	val := "1"
	desk := "2"
	url := "3"
	tests := []struct {
		name    string
		g       *GithubAdapter
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
		name:    "with cache",
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
		args:    args{ctx: context.Background(), username: ""},
		want:    []Item{},
		wantErr: true,
		list:    []*github.Gist{{ID: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 100}},
		errList: nil,
	}, {
		name:    "with error",
		g:       g,
		args:    args{username: ""},
		want:    []Item{},
		wantErr: true,
		list:    []*github.Gist{{ID: &val, Description: &desk, HTMLURL: &url}},
		resp:    &github.Response{Response: &http.Response{StatusCode: 200}},
		errList: errors.New("ctx can not be nil"),
	}}
	for _, tt := range tests {
		if tt.name != "with cache" {
			gisMock.On("List", tt.args.ctx, tt.args.username, mock.AnythingOfType("*github.GistListOptions")).Return(tt.list, tt.resp, tt.errList)
		}
		got, err := tt.g.GetGists(tt.args.ctx, tt.args.username)
		if (err != nil) != tt.wantErr {
			t.Errorf("GithubAdapter.GetGists() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.Equal(t, tt.want, got)
	}
}

func TestList(t *testing.T) {
	gt := github.NewClient(&http.Client{})
	g := NewGithub(gt)
	_, _, err := g.Gister.List(nil, "sad", nil)
	if err == nil {
		t.Errorf("should be error")
	}
	_, _, err = g.Reposer.List(nil, "sad", nil)
	if err == nil {
		t.Errorf("should be error")
	}
}
