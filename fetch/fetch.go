package fetch

import (
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

var cuitere = regexp.MustCompile(`cuite\.v(\d+)\..*`)

// Client is the struct holding our data and client to fetch data on Soundcloud
type Client struct {
	ids      []uint64
	clientID string
	c        *soundcloud.Api

	Playlists     []*soundcloud.Playlist
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
		} else {
			return &RichPlaylist{p, 0}
		}
	}
	return &RichPlaylist{p, 0}
}

// NewClient returns a new Clinet
func NewClient(clientID string, ids []uint64) *Client {
	return &Client{
		ids:      ids,
		clientID: clientID,
		c: &soundcloud.Api{
			ClientId: clientID,
		},
	}
}

// Fetch will fetch and store the playlists and tracks into the client
func (c *Client) Fetch() error {
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
	for i := 0; i < len(c.RichPlaylists); i++ {
		c.Playlists[i] = c.RichPlaylists[i].Playlist
	}
	c.Tracks = []*soundcloud.Track{}
	for _, p := range c.Playlists {
		c.Tracks = append(c.Tracks, p.Tracks...)
	}
	return nil
}

// Find will fetch and find the matching playlists of a given user id
func (c *Client) Find(id uint64) ([]*RichPlaylist, error) {
	var out []*RichPlaylist
	var pl []*soundcloud.Playlist
	var err error

	if pl, err = c.c.User(id).Playlists(url.Values{}); err != nil {
		return out, errors.Wrapf(err, "fetch playlists for %d", id)
	}
	for _, p := range pl {
		if strings.HasPrefix(p.Title, "cuite.v") {
			out = append(out, NewRichPlaylist(p))
		}
	}
	return out, nil
}
