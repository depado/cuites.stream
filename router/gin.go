package router

import (
	"net/http"

	"github.com/Depado/cuitesite/models"
	"github.com/gin-gonic/gin"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

// GinRouter is a simple router that embeds the data
type GinRouter struct {
	Cuites []*soundcloud.Playlist
}

// GetPlaylists will return the playlists in JSON format
func (r GinRouter) GetPlaylists(c *gin.Context) {
	c.JSON(http.StatusOK, models.FormatPlaylists(r.Cuites, false))
}
