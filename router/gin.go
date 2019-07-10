package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

// GinRouter is a simple router that embeds the data
type GinRouter struct {
	Playlists    []*soundcloud.Playlist
	Tracks       []*soundcloud.Track
	PlaylistsMap *sync.Map
}

// AddRoutes will bind routes to an existing engine
func (r *GinRouter) AddRoutes(e *gin.Engine) {
	e.GET("/playlists", r.GetPlaylists)
	e.GET("/tracks", r.GetAllTracks)
	e.GET("/playlist/:id/tracks", r.GetPlaylistTracks)
}
