package fetch

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dghubble/sling"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

var cuitere = regexp.MustCompile(`cuite\.v(\d+)\..*`)

// Client is the struct holding our data and client to fetch data on Soundcloud
type Client struct {
	ids      []int
	clientID string
	c        *soundcloud.Api
	hc       *http.Client

	Playlists     []*soundcloud.Playlist
	PlaylistsMap  sync.Map
	RichPlaylists []*RichPlaylist
	Tracks        []*soundcloud.Track
}

// RichPlaylist is just a playlist with its number parsed
type RichPlaylist struct {
	*soundcloud.Playlist
	Number int
}

// NewRichPlaylist will parse the playlist's title to fill the number
func NewRichPlaylist(p *soundcloud.Playlist) *RichPlaylist {
	m := cuitere.FindStringSubmatch(p.Title)
	if len(m) >= 2 {
		if i, err := strconv.Atoi(m[1]); err == nil {
			return &RichPlaylist{p, i}
		}
	}
	return &RichPlaylist{p, 0}
}

// NewClient returns a new Clinet
func NewClient(clientID string, ids []int) *Client {
	return &Client{
		ids:      ids,
		clientID: clientID,
		c: &soundcloud.Api{
			ClientId: clientID,
		},
		hc: &http.Client{
			Timeout: time.Duration(5 * time.Second),
		},
	}
}

// Fetch will fetch and store the playlists and tracks into the client
func (c *Client) Fetch() error {
	logrus.Info("Fetching and parsing playlists")
	start := time.Now()
	c.RichPlaylists = []*RichPlaylist{}
	for _, id := range c.ids {
		p, err := c.Find(id)
		if err != nil {
			return errors.Wrap(err, "fetch")
		}
		c.RichPlaylists = append(c.RichPlaylists, p...)
	}
	sort.Slice(c.RichPlaylists, func(i, j int) bool {
		return c.RichPlaylists[i].Number > c.RichPlaylists[j].Number
	})
	c.Playlists = make([]*soundcloud.Playlist, len(c.RichPlaylists))
	c.Tracks = []*soundcloud.Track{}
	for i := 0; i < len(c.RichPlaylists); i++ {
		c.Playlists[i] = c.RichPlaylists[i].Playlist
		c.PlaylistsMap.Store(c.RichPlaylists[i].Id, c.RichPlaylists[i].Playlist)
		c.Tracks = append(c.Tracks, c.RichPlaylists[i].Playlist.Tracks...)
	}
	logrus.WithField("took", time.Since(start)).Info("Done")
	return nil
}

// Find will fetch and find the matching playlists of a given user id
func (c *Client) Find(id int) ([]*RichPlaylist, error) {
	var out []*RichPlaylist
	var pl []*soundcloud.Playlist
	var err error

	if pl, err = c.c.User(uint64(id)).Playlists(url.Values{}); err != nil {
		return out, errors.Wrapf(err, "fetch playlists for %d", id)
	}
	for _, p := range pl {
		if strings.HasPrefix(p.Title, "cuite.v") {
			out = append(out, NewRichPlaylist(p))
		}
	}
	return out, nil
}

// Stream will retrieve a streamable content respecting SoundCloud's workflow
func (c *Client) Stream(id string) ([]byte, error) {
	var err error
	var resp *http.Response
	var req *http.Request

	if req, err = sling.New().
		Get("https://api.soundcloud.com/i1/tracks/" + id + "/streams").
		QueryStruct(struct {
			ClientID string `url:"client_id"`
		}{ClientID: c.clientID}).
		Request(); err != nil {
		return nil, err
	}
	if resp, err = c.hc.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bb, nil
}
