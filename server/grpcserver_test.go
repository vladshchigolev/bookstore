package server

import (
	"bookstore/pkg/api"
	"context"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGRPCServer_GetBooks(t *testing.T) {
	config := NewConfig()
	_, err := toml.DecodeFile("C:\\Users\\amrgF\\GolandProjects\\bookstore\\configs\\grpcserver.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	srv := New(config)
	if err := srv.ConfigureLogger(); err != nil {
		log.Fatal(err)
	}
	if err := srv.ConfigureDatabase(); err != nil {
		log.Fatal(err)
	}

	author := api.Author{Name: "Lisa Urry"}
	ctx := context.Background()
	_, err = srv.GetBooks(ctx, &author)
	//if bookSet.Books !=
	assert.NoError(t, err)
}

func TestGRPCServer_GetAuthors(t *testing.T) {
	config := NewConfig()
	_, err := toml.DecodeFile("C:\\Users\\amrgF\\GolandProjects\\bookstore\\configs\\grpcserver.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	srv := New(config)
	if err := srv.ConfigureLogger(); err != nil {
		log.Fatal(err)
	}
	if err := srv.ConfigureDatabase(); err != nil {
		log.Fatal(err)
	}

	title := api.Title{Title: "Campbell Biology (Campbell Biology Series)"}
	ctx := context.Background()
	_, err = srv.GetAuthors(ctx, &title)
	//if bookSet.Books !=
	assert.NoError(t, err)
}
