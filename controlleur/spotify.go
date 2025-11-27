package controlleur

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
    tokenCache  string
    tokenExpiry time.Time
)

// -----------------------------------------------------
// 🔐 Récupération du token Spotify (avec gestion erreur)
// -----------------------------------------------------
func spotifyToken() (string, error) {

    clientID := os.Getenv("SPOTIFY_CLIENT_ID")
    secret := os.Getenv("SPOTIFY_CLIENT_SECRET")

    if clientID == "" || secret == "" {
        return "", fmt.Errorf("spotify: SPOTIFY_CLIENT_ID ou SPOTIFY_CLIENT_SECRET vide")
    }

    if tokenCache != "" && time.Now().Before(tokenExpiry) {
        return tokenCache, nil
    }

    data := url.Values{}
    data.Set("grant_type", "client_credentials")

    req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token",
        strings.NewReader(data.Encode()))

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization",
        "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+secret)))

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        b, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("spotify token error %d: %s", resp.StatusCode, string(b))
    }

    var auth struct {
        Token string `json:"access_token"`
        Exp   int    `json:"expires_in"`
    }

    json.NewDecoder(resp.Body).Decode(&auth)

    tokenCache = auth.Token
    tokenExpiry = time.Now().Add(time.Duration(auth.Exp) * time.Second)

    return tokenCache, nil
}

// -----------------------------------------------------
// 🎵 Albums Damso
// -----------------------------------------------------
func GetDamsoAlbums() ([]Album, error) {

    token, err := spotifyToken()
    if err != nil {
        return nil, err
    }

    req, _ := http.NewRequest("GET",
        "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums?include_groups=album&market=FR",
        nil)

    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        b, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("spotify albums error %d: %s", resp.StatusCode, string(b))
    }

    var result struct {
        Items []Album `json:"items"`
    }

    json.NewDecoder(resp.Body).Decode(&result)
    return result.Items, nil
}

// -----------------------------------------------------
// 🎧 Track Laylow – Maladresse
// -----------------------------------------------------
func GetLaylowTrack() (Track, error) {

    token, err := spotifyToken()
    if err != nil {
        return Track{}, err
    }

    req, _ := http.NewRequest("GET",
        "https://api.spotify.com/v1/tracks/67Pf31pl0PfjBfUmvYNDCL?market=FR",
        nil)

    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return Track{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        b, _ := io.ReadAll(resp.Body)
        return Track{}, fmt.Errorf("spotify track error %d: %s", resp.StatusCode, string(b))
    }

    var track Track
    json.NewDecoder(resp.Body).Decode(&track)

    return track, nil
}
