package detect

import (
	"strings"

	"github.com/dvrg/useragent"
)

func isiOS(os, platform string) bool {
	return contains(platform, "ipad") ||
		contains(platform, "ipod") ||
		contains(platform, "iphone")
}

func isWindowsPhone(userAgent, os, platform string, isMobile bool) bool {
	return ((contains(platform, "windows") ||
		contains(os, "windows")) && isMobile) ||
		contains(strings.ToLower(userAgent), "iemobile")
}

func isAndroid(os, platform string) bool {
	return contains(platform, "android") ||
		contains(os, "android")
}

func contains(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}

// GetPlatform returns the platform of the device with the given User-Agent
func GetPlatform(userAgent string) Platform {
	var (
		ua       = useragent.New(userAgent)
		os       = strings.ToLower(ua.OS())
		platform = strings.ToLower(ua.Platform())
	)

	switch true {
	case isiOS(os, platform):
		return IOS
	case isAndroid(os, platform):
		return ANDROID
	case isWindowsPhone(userAgent, os, platform, ua.Mobile()):
		return WINDOWSPHONE
	case ua.Mobile():
		return MOBILE
	default:
		return DESKTOP
	}
}

// IsMobile returns true if the visitor is using a mobile device
func IsMobile(ua string) bool {
	switch GetPlatform(ua) {
	case IOS, ANDROID, WINDOWSPHONE, MOBILE:
		return true
	}
	return false
}
