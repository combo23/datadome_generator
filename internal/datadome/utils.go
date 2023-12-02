package datadome

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/combo23/datadome_generator/internal/constants"
)

func calculateMcmr(mousemove, click int) int {
	if mousemove == 0 {
		return -1
	}
	return click / mousemove
}

func calculateMmsr(mousemove, scroll int) int {
	if scroll == 0 {
		return -1
	}
	return mousemove / scroll
}

func randomBoolean(rand.Rand) bool {
	return rand.Intn(2) == 0
}

// idk why there is "Other" option but this is in js script
func getEvavalue(ua string) (int, error) {
	if strings.Contains(ua, "Chrome") || strings.Contains(ua, "Opera") || strings.Contains(ua, "Other") {
		return 37, nil
	} else if strings.Contains(ua, "Internet Explorer") || strings.Contains(ua, "Other") {
		return 39, nil
	} else if strings.Contains(ua, "Firefox") || strings.Contains(ua, "Safari") || strings.Contains(ua, "Other") {
		return 33, nil
	}
	return 0, fmt.Errorf("unknown user agent")
}

func extractPathFromURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	decodedPath, err := url.QueryUnescape(parsedURL.Path)
	if err != nil {
		return "", err
	}

	return decodedPath, nil
}

func getDDInfo(host string) (string, string) {
	for x := range constants.SupportedSites {
		if strings.Contains(host, x) {
			return constants.SupportedSites[x].DDK, constants.SupportedSites[x].DDVersion
		}
	}
	return "", ""
}
