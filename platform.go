package detect

// PlatformType is the kind of platform of the device.
type PlatformType byte

const (
	// Desktop operating system
	Desktop PlatformType = 1 << iota
	// Mobile operating system (android, ios, WP, ...)
	Mobile
	// Android operating system
	Android
	// IOs operating system (iPhone, iPad, iPod, ...)
	IOs
	// WindowsPhone operating system
	WindowsPhone
)
