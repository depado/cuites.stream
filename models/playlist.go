package models

import (
	"github.com/yanatan16/golang-soundcloud/soundcloud"
)

// FormatTracks will format a track slice
func FormatTracks(t []*soundcloud.Track) []*Track {
	tt := make([]*Track, len(t))
	for i := 0; i < len(t); i++ {
		tt[i] = FormatTrack(t[i])
	}
	return tt
}

// FormatTrack will format a track
func FormatTrack(t *soundcloud.Track) *Track {
	tt := &Track{
		ID:               t.Id,
		CreatedAt:        t.CreatedAt,
		Title:            t.Title,
		ArtworkURL:       t.ArtworkUrl,
		AssetData:        t.AssetData,
		ArtworkData:      t.ArtworkData,
		Description:      t.Description,
		Duration:         t.Duration,
		Genre:            t.Genre,
		TagList:          t.TagList,
		ReleaseDay:       t.ReleaseDay,
		ReleaseMonth:     t.ReleaseMonth,
		ReleaseYear:      t.ReleaseYear,
		CommentCount:     t.CommentCount,
		DownloadCount:    t.DownloadCount,
		PlaybackCount:    t.PlaybackCount,
		FavoritingsCount: t.FavoritingsCount,
	}

	if t.SubUri != nil {
		tt.SubURI = &SubURI{URI: t.SubUri.Uri}
	}
	if t.SubUser != nil {
		tt.SubUser = &SubUser{User: FormatUser(t.SubUser.User)}
	}
	if t.SubPermalink != nil {
		tt.SubPermalink = &SubPermalink{PermalinkURL: t.SubPermalink.PermalinkUrl}
	}
	return tt
}

// FormatPlaylists will format a playlist slice
func FormatPlaylists(p []*soundcloud.Playlist, tracks bool) []*Playlist {
	pp := make([]*Playlist, len(p))
	for i := 0; i < len(p); i++ {
		pp[i] = FormatPlaylist(p[i], tracks)
	}
	return pp
}

// FormatPlaylist formats a soundcloud playlist to a minimal version
func FormatPlaylist(p *soundcloud.Playlist, tracks bool) *Playlist {
	pp := &Playlist{
		ID:           p.Id,
		CreatedAt:    p.CreatedAt,
		Title:        p.Title,
		ArtworkURL:   p.ArtworkUrl,
		Description:  p.Description,
		Duration:     p.Duration,
		Genre:        p.Genre,
		TagList:      p.TagList,
		ReleaseDay:   p.ReleaseDay,
		ReleaseMonth: p.ReleaseMonth,
		ReleaseYear:  p.ReleaseYear,
		PlaylistType: p.PlaylistType,
	}
	if tracks {
		pp.Tracks = FormatTracks(p.Tracks)
	}
	if p.SubUri != nil {
		pp.SubURI = &SubURI{URI: p.SubUri.Uri}
	}
	if p.SubUser != nil {
		pp.SubUser = &SubUser{User: FormatUser(p.SubUser.User)}
	}
	if p.SubPermalink != nil {
		pp.SubPermalink = &SubPermalink{PermalinkURL: p.SubPermalink.PermalinkUrl}
	}
	return pp
}

// FormatUser formats a soundcloud user to a minimal version
func FormatUser(u *soundcloud.User) *User {
	uu := &User{
		ID:         u.Id,
		Username:   u.Username,
		AvatarURL:  u.AvatarUrl,
		AvatarData: u.AvatarData,
	}
	if u.SubUri != nil {
		uu.SubURI = &SubURI{URI: u.SubUri.Uri}
	}
	if u.SubPermalink != nil {
		uu.SubPermalink = &SubPermalink{PermalinkURL: u.SubPermalink.PermalinkUrl}
	}
	return uu
}

// User represents a user
type User struct {
	ID uint64 `json:"id,omitempty"`

	*SubURI
	*SubPermalink

	Username   string `json:"username,omitempty"`
	AvatarURL  string `json:"avatar_url,omitempty"`
	AvatarData []byte `json:"avatar_data,omitempty"`
}

// Track represents a single track
type Track struct {
	ID uint64 `json:"id,omitempty"`

	*SubURI
	*SubUser
	*SubPermalink

	CreatedAt        string `json:"created_at,omitempty"`
	Title            string `json:"title,omitempty"`
	ArtworkURL       string `json:"artwork_url,omitempty"`
	Description      string `json:"description,omitempty"`
	Duration         uint64 `json:"duration,omitempty"`
	Genre            string `json:"genre,omitempty"`
	TagList          string `json:"tag_list,omitempty"`
	ReleaseDay       uint   `json:"release_day,omitempty"`
	ReleaseMonth     uint   `json:"release_month,omitempty"`
	ReleaseYear      uint   `json:"release_year,omitempty"`
	CommentCount     uint64 `json:"comment_count,omitempty"`
	DownloadCount    uint64 `json:"download_count,omitempty"`
	PlaybackCount    uint64 `json:"playback_count,omitempty"`
	FavoritingsCount uint64 `json:"favoritings_count,omitempty"`
	AssetData        []byte `json:"asset_data,omitempty"`
	ArtworkData      []byte `json:"artwork_data,omitempty"`
}

// Playlist represents a single playlist
type Playlist struct {
	ID uint64 `json:"id,omitempty"`

	*SubURI
	*SubUser
	*SubPermalink

	CreatedAt    string   `json:"created_at,omitempty"`
	Title        string   `json:"title,omitempty"`
	ArtworkURL   string   `json:"artwork_url,omitempty"`
	Description  string   `json:"description,omitempty"`
	Duration     uint64   `json:"duration,omitempty"`
	Genre        string   `json:"genre,omitempty"`
	TagList      string   `json:"tag_list,omitempty"`
	ReleaseDay   uint     `json:"release_day,omitempty"`
	ReleaseMonth uint     `json:"release_month,omitempty"`
	ReleaseYear  uint     `json:"release_year,omitempty"`
	PlaylistType string   `json:"playlist_type,omitempty"`
	Tracks       []*Track `json:"tracks,omitempty"`
}

// Comment represents a single comment
type Comment struct {
	ID uint64 `json:"id,omitempty"`

	*SubURI
	*SubUser

	CreatedAt string `json:"created_at,omitempty"`
	Body      string `json:"body,omitempty"`
	// Time in milliseconds
	Timestamp uint64 `json:"timestamp,omitempty"`
	TrackID   uint64 `json:"track_id,omitempty"`
}

// SubURI is a sub structure containing an URI
type SubURI struct {
	URI string `json:"uri,omitempty"`
}

// SubPermalink is a substructure embedded in various structs
type SubPermalink struct {
	PermalinkURL string `json:"permalink_url"`
}

// SubUser is a substructure embedded in various other structs
type SubUser struct {
	UserID uint64 `json:"user_id,omitempty"`
	User   *User  `json:"user,omitempty"`
}
