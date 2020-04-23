package service

import (
	"net/http"

	"github.com/varmamsp/cello/web"
)

const (
	LOAD_SESSION        = "load_session"
	END_SESSION         = "end_session"
	ADD_TO_PLAYLIST     = "add_to_playlist"
	CREATE_PLAYLIST     = "create_playlist"
	EDIT_PLAYLIST       = "edit_playlist"
	DELETE_PLAYLIST     = "delete_playlist"
	GET_PLAYBACKS       = "get_playbacks"
	PLAYBACK_SYNC       = "playback_sync"
	SUBSCRIBE_PODCAST   = "subscribe_podcast"
	UNSUBSCRIBE_PODCAST = "unsubscribe_podcast"

	ACTION_ADD_EPISODE       = "add_episode"
	ACTION_REMOVE_EPISODE    = "remove_episode"
	ACTION_PLAYBACK_BEGIN    = "playback_begin"
	ACTION_PLAYBACK_PROGRESS = "playback_progress"
)

func RootHandler(c *web.Context, w http.ResponseWriter, req *http.Request) {
	if c.RequireEndpoint(); c.Err != nil {
		return
	}

	switch c.Params.Endpoint {
	case LOAD_SESSION:
		loadSession(c, w, req)

	case END_SESSION:
		endSession(c, w, req)

	case ADD_TO_PLAYLIST:
		addToPlaylist(c, w, req)

	case CREATE_PLAYLIST:
		createPlaylist(c, w, req)

	case EDIT_PLAYLIST:
		if c.RequireAction(); c.Err != nil {
			return
		} else if action := c.Params.Action; action == ACTION_ADD_EPISODE {
			addEpisodeToPlaylist(c, w, req)
		} else if action == ACTION_REMOVE_EPISODE {
			removeEpisodeFromPlaylist(c, w, req)
		} else {
			c.SetInvalidQueryParam("action")
		}

	case DELETE_PLAYLIST:
		deletePlaylist(c, w, req)

	case GET_PLAYBACKS:
		getPlaybacks(c, w, req)

	case PLAYBACK_SYNC:
		if c.RequireAction(); c.Err != nil {
			return
		} else if action := c.Params.Action; action == ACTION_PLAYBACK_BEGIN {
			syncPlaybackBegin(c, w, req)
		} else if action == ACTION_PLAYBACK_PROGRESS {
			syncPlaybackProgress(c, w, req)
		} else {
			c.SetInvalidQueryParam("action")
		}

	case SUBSCRIBE_PODCAST:
		subscribeToPodcast(c, w, req)

	case UNSUBSCRIBE_PODCAST:
		unsubscribeToPodcast(c, w, req)

	default:
		c.SetInvalidQueryParam("endpoint")
	}
}
