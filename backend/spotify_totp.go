package backend

import (
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	spotifyTOTPVersion = 61
)

func generateSpotifyTOTP(now time.Time) (string, int, error) {
	spotifyTOTPSecret := os.Getenv("SPOTIFY_TOTP_SECRET")
	if spotifyTOTPSecret == "" {
		return "", 0, fmt.Errorf("SPOTIFY_TOTP_SECRET environment variable not set")
	}
	key, err := otp.NewKeyFromURL(fmt.Sprintf("otpauth://totp/secret?secret=%s", spotifyTOTPSecret))
	if err != nil {
		return "", 0, err
	}

	code, err := totp.GenerateCode(key.Secret(), now)
	if err != nil {
		return "", 0, err
	}

	return code, spotifyTOTPVersion, nil
}
