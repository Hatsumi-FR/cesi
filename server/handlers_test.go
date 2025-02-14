package server

import (
	"cesi/adapters/inmemory"
	"testing"
)

func TestServer_CreatePlayerHandler(t *testing.T) {
	repo := &inmemory.Repository{}
	s := NewServer(repo)
	s.CreatePlayerHandler(nil, nil)
}
