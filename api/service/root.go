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
	c.RequireEndpoint()
	if c.Err != nil {
		return
	}

	switch c.Params.Endpoint {
	case LOAD_SESSION:
		loadSession(c, w, req)

	case END_SESSION:
		endSession(c, w, req)

	case ADD_TO_PLAYLIST:
		addEpisodeToPlaylist(c, w, req)

	case CREATE_PLAYLIST:
		createPlaylist(c, w, req)

	case EDIT_PLAYLIST:
		c.RequireAction()
		if c.Err == nil {
			switch c.Params.Action {
			case ACTION_ADD_EPISODE:
				addEpisodeToPlaylist(c, w, req)

			case ACTION_REMOVE_EPISODE:
				removeEpisodeFromPlaylist(c, w, req)

			default:
				c.SetInvalidQueryParam("action")
			}
		}

	case DELETE_PLAYLIST:
		deletePlaylist(c, w, req)

	case GET_PLAYBACKS:
		getPlaybacks(c, w, req)

	case PLAYBACK_SYNC:
		c.RequireSession().RequireAction().RequireBody(req)
		if c.Err == nil {
			switch c.Params.Action {
			case ACTION_PLAYBACK_BEGIN:

			case ACTION_PLAYBACK_PROGRESS:

			default:
				c.SetInvalidQueryParam("endpoint")
			}
		}

	case SUBSCRIBE_PODCAST:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			subscribeToPodcast(c, w, req)
		}

	case UNSUBSCRIBE_PODCAST:
		c.RequireSession().RequireBody(req)
		if c.Err == nil {
			unsubscribeToPodcast(c, w, req)
		}

	default:
		c.SetInvalidQueryParam("endpoint")
	}
}
