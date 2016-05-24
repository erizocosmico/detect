package detect

import (
	"strings"

	"github.com/mvader/useragent"
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

// Platform returns the platform of the device with the given User-Agent.
func Platform(userAgent string) PlatformType {
	var (
		ua       = useragent.New(userAgent)
		os       = strings.ToLower(ua.OS())
		platform = strings.ToLower(ua.Platform())
	)

	switch true {
	case isiOS(os, platform):
		return IOs
	case isAndroid(os, platform):
		return Android
	case isWindowsPhone(userAgent, os, platform, ua.Mobile()):
		return WindowsPhone
	case ua.Mobile():
		return Mobile
	default:
		return Desktop
	}
}

// IsMobile returns true if the visitor is using a mobile device
func IsMobile(ua string) bool {
	switch Platform(ua) {
	case IOs, Android, WindowsPhone, Mobile:
		return true
	}
	return false
}
