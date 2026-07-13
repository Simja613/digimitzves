package registry

type Device struct {
	ID string

	Parent string

	FriendlyName string

	Present bool

	Configured bool

	Ignored bool
}
