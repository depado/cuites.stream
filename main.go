package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Depado/cuitesite/cmd"
	"github.com/Depado/cuitesite/fetch"
	"github.com/Depado/cuitesite/router"
)

// Build number and versions injected at compile time
var (
	Version = "unknown"
	Build   = "unknown"
)

// Main command that will be run when no other command is provided on the
// command-line
var rootc = &cobra.Command{
	Use:   "cuitesite <options>",
	Short: "Cuitesite backend",
	Long:  "Backend app that will aggregate playlists",
	Run: func(cmd *cobra.Command, args []string) {
		fc := fetch.NewClient(viper.GetString("client_id"), []uint64{17771323, 93734268, 20836701, 153939520, 39713634})
		if err := fc.Fetch(); err != nil {
			logrus.WithError(err).Fatal("Unable to fetch content from Soundcloud")
		}
		gr := router.GinRouter{Playlists: fc.Playlists, Tracks: fc.Tracks}

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
