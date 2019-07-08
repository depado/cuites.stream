package router

import (
	"net/http"

	"github.com/Depado/cuitesite/models"
	"github.com/gin-gonic/gin"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

// GinRouter is a simple router that embeds the data
type GinRouter struct {
	Playlists []*soundcloud.Playlist
	Tracks    []*soundcloud.Track
}

// GetPlaylists will return the playlists in JSON format
func (r GinRouter) GetPlaylists(c *gin.Context) {
	c.JSON(http.StatusOK, models.FormatPlaylists(r.Playlists, false))
}

// GetAllTracks will return all the tracks in JSON format
func (r GinRouter) GetAllTracks(c *gin.Context) {
	c.JSON(http.StatusOK, models.FormatTracks(r.Tracks))
}
