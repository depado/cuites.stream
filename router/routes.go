package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yanatan16/golang-soundcloud/soundcloud"

	"github.com/Depado/cuites.stream/models"
)

// GetPlaylists will return the playlists in JSON format
func (r *GinRouter) GetPlaylists(c *gin.Context) {
	c.JSON(http.StatusOK, models.FormatPlaylists(r.Playlists, false))
}

// GetAllTracks will return all the tracks in JSON format
func (r *GinRouter) GetAllTracks(c *gin.Context) {
	c.JSON(http.StatusOK, models.FormatTracks(r.Tracks))
}

// GetPlaylistTracks will retrieve the playlist tracks
func (r *GinRouter) GetPlaylistTracks(c *gin.Context) {
	idp := c.Param("id")
	id, err := strconv.Atoi(idp)
	if err != nil {
		logrus.WithError(err).Error("unable to parse playlist ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	p, ok := r.PlaylistsMap.Load(uint64(id))
	if !ok {
		logrus.WithField("id", id).Error("couldn't find playlist")
		c.JSON(http.StatusNotFound, gin.H{"id": id, "error": "not found"})
		return
	}
	c.JSON(http.StatusOK, models.FormatTracks(p.(*soundcloud.Playlist).Tracks))
}
