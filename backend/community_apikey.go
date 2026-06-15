package backend

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	communityAPIKeyOnce sync.Once
	communityAPIKey     string
	communityAPIKeyErr  error
)

func getCommunityAPIKey() (string, error) {
	communityAPIKeyOnce.Do(func() {
		communityAPIKey = os.Getenv("SPOTIFLAC_COMMUNITY_API_KEY")
		if communityAPIKey == "" {
			communityAPIKeyErr = fmt.Errorf("SPOTIFLAC_COMMUNITY_API_KEY environment variable not set")
		}
	})

	if communityAPIKeyErr != nil {
		return "", communityAPIKeyErr
	}
	return communityAPIKey, nil
}

func communityUserAgent() string {
	version := strings.TrimSpace(AppVersion)
	if version == "" || version == "Unknown" {
		return "SpotiFLAC"
	}
	return "SpotiFLAC/" + version
}

func setCommunityRequestHeaders(req *http.Request) error {
	apiKey, err := getCommunityAPIKey()
	if err != nil {
		return fmt.Errorf("failed to prepare community API key: %w", err)
	}
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("User-Agent", communityUserAgent())
	return nil
}
