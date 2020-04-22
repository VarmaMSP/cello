package service

import (
	"net/http"

	"github.com/varmamsp/cello/web"
)

func addToPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {}

func createPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {}

func deletePlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {}

func addEpisodeToPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {}

func removeEpisodeFromPlaylist(c *web.Context, w http.ResponseWriter, req *http.Request) {}
