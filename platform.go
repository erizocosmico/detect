package detect

type Platform byte

const (
	// Desktop operating system
	DESKTOP Platform = iota
	// Mobile operating system (android, ios, WP, ...)
	MOBILE
	// Android operating system
	ANDROID
	// iOS operating system (iPhone, iPad, iPod, ...)
	IOS
	// Windows Phone operating system
	WINDOWSPHONE
)
