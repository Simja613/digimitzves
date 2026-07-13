package registry

import "testing"

func TestRegistryAdd(t *testing.T) {

	r := Registry{}

	r.Add(
		Device{
			ID: "relay#1",
		},
	)

	if !r.Exists("relay#1") {
		t.Fatal("device not found")
	}

}

func TestKnownDevice(t *testing.T) {

	r := Registry{}

	device := Device{
		ID: "relay#1",
	}

	r.Add(device)

	found, ok := r.Find("relay#1")

	if !ok {
		t.Fatal("device should exist")
	}

	if found.ID != device.ID {
		t.Fatal("wrong device returned")
	}

}

func TestUnknownDevice(t *testing.T) {

	r := Registry{}

	found, ok := r.Find("relay#1")

	if ok {
		t.Fatal("device should not exist")
	}

	if found != nil {
		t.Fatal("expected nil device")
	}

}
