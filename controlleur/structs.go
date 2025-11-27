package controlleur

// Image Spotify
type Image struct {
    URL string `json:"url"`
}

// Album Spotify
type Album struct {
    Name        string  `json:"name"`
    Images      []Image `json:"images"`
    ReleaseDate string  `json:"release_date"`
    TotalTracks int     `json:"total_tracks"`
}

// Track Spotify
type Track struct {
    Name   string `json:"name"`

    Album struct {
        Name        string  `json:"name"`
        Images      []Image `json:"images"`
        ReleaseDate string  `json:"release_date"`
    } `json:"album"`

    Artists []struct {
        Name string `json:"name"`
    } `json:"artists"`

    ExternalURLs map[string]string `json:"external_urls"`
}
