package detect

// PlatformType is the kind of platform of the device.
type PlatformType byte

const (
	// DESKTOP operating system
	DESKTOP PlatformType = 1 << iota
	// MOBILE operating system (android, ios, WP, ...)
	MOBILE
	// ANDROID operating system
	ANDROID
	// IOS operating system (iPhone, iPad, iPod, ...)
	IOS
	// WINDOWSPHONE operating system
	WINDOWSPHONE
)
