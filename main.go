package main

import (
	"log"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Depado/cuites.stream/cmd"
	"github.com/Depado/cuites.stream/fetch"
	"github.com/Depado/cuites.stream/infra"
	"github.com/Depado/cuites.stream/router"
)

// Build number and versions injected at compile time
var (
	Version = "unknown"
	Build   = "unknown"
)

// Main command that will be run when no other command is provided on the
// command-line
var rootc = &cobra.Command{
	Use:   "cuites.stream <options>",
	Short: "cuites.stream backend",
	Long:  "Backend app that will aggregate playlists",
	Run:   func(cmd *cobra.Command, args []string) { run() },
}

func run() {
	strids := viper.GetStringSlice("user_ids")
	ids := make([]int, len(strids))
	for i := 0; i < len(strids); i++ {
		var err error
		if ids[i], err = strconv.Atoi(strids[i]); err != nil {
			logrus.WithError(err).Fatal("not an integer ID")
		}
	}
	s := infra.NewServer(
		viper.GetString("server.host"),
		viper.GetInt("server.port"),
		viper.GetString("server.mode"),
		infra.NewCorsConfig(
			viper.GetBool("server.cors.disabled"),
			viper.GetBool("server.cors.all"),
			viper.GetStringSlice("server.cors.origins"),
			viper.GetStringSlice("server.cors.methods"),
			viper.GetStringSlice("server.cors.headers"),
			viper.GetStringSlice("server.cors.expose"),
		),
	)
	fc := fetch.NewClient(
		viper.GetString("client_id"),
		ids,
	)
	if err := fc.Fetch(); err != nil {
		logrus.WithError(err).Fatal("Unable to fetch content from Soundcloud")
	}
	gr := &router.GinRouter{
		Playlists:    fc.Playlists,
		Tracks:       fc.Tracks,
		PlaylistsMap: &fc.PlaylistsMap,
		Fetcher:      fc,
	}
	gr.AddRoutes(s.Router)
	s.Start()
}

func main() {
	// Initialize Cobra and Viper
	cobra.OnInitialize(cmd.Initialize)
	cmd.AddAllFlags(rootc)

	// Run the command
	if err := rootc.Execute(); err != nil {
		log.Fatal("Couldn't start program:", err)
	}
}
