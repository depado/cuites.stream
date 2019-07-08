package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yanatan16/golang-soundcloud/soundcloud"

	"github.com/Depado/cuitesite/cmd"
	"github.com/Depado/cuitesite/router"
)

// Build number and versions injected at compile time
var (
	Version = "unknown"
	Build   = "unknown"
)

func find(c *soundcloud.Api, id uint64) ([]*soundcloud.Playlist, error) {
	var out, pl []*soundcloud.Playlist
	var err error

	if pl, err = c.User(id).Playlists(url.Values{}); err != nil {
		return out, errors.Wrapf(err, "fetch playlists for %d", id)
	}
	for _, p := range pl {
		if strings.HasPrefix(p.Title, "cuite.v") {
			out = append(out, p)
		}
	}
	return out, nil
}

func fetch(c *soundcloud.Api) ([]*soundcloud.Playlist, error) {
	var out []*soundcloud.Playlist

	ids := []uint64{17771323, 93734268, 20836701, 153939520, 39713634}
	for _, id := range ids {
		p, err := find(c, id)
		if err != nil {
			return out, errors.Wrap(err, "fetch")
		}
		out = append(out, p...)
	}
	return out, nil
}

func tracks(pp []*soundcloud.Playlist) []*soundcloud.Track {
	tracks := []*soundcloud.Track{}
	for _, p := range pp {
		tracks = append(tracks, p.Tracks...)
	}
	return tracks
}

// Main command that will be run when no other command is provided on the
// command-line
var rootc = &cobra.Command{
	Use:   "cuitesite <options>",
	Short: "Cuitesite backend",
	Long:  "Backend app that will aggregate playlists",
	Run: func(cmd *cobra.Command, args []string) {
		c := &soundcloud.Api{
			ClientId: viper.GetString("client_id"),
		}

		pl, err := fetch(c)
		if err != nil {
			logrus.WithError(err).Fatal("Unable to fetch playlists")
		}
		t := tracks(pl)
		gr := router.GinRouter{Playlists: pl, Tracks: t}

		r := gin.Default()
		r.Use(cors.New(cors.Config{
			AllowCredentials: true,
			MaxAge:           50 * time.Second,
			AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTION", "PATCH"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			AllowAllOrigins:  true,
		}))

		r.GET("/playlists", gr.GetPlaylists)
		r.GET("/tracks", gr.GetAllTracks)
		r.Run("127.0.0.1:8081")
	},
}

// Version command that will display the build number and version (if any)
var versionc = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run:   func(c *cobra.Command, args []string) { fmt.Printf("Build: %s\nVersion: %s\n", Build, Version) },
}

func main() {
	// Initialize Cobra and Viper
	cobra.OnInitialize(cmd.Initialize)
	cmd.AddLoggerFlags(rootc)
	cmd.AddGlobalFlags(rootc)

	// Run the command
	if err := rootc.Execute(); err != nil {
		log.Fatal("Couldn't start program:", err)
	}
}
