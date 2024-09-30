package reverse

import (
	"os"
)

// ReadBannerFile reads the content of a banner file specified by the filepath argument and returns it as a string.
func ReadBannerFile(filepath string) (string, error) {
	bannerFile, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(bannerFile), nil
}
