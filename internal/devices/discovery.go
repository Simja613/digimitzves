package devices

type DiscoveredDevice struct {
	Parent string

	Channel int

	FriendlyName string
}

func Discover() ([]DiscoveredDevice, error) {

	return []DiscoveredDevice{}, nil

}
