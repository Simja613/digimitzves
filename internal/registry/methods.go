package registry

func (r *Registry) Exists(id string) bool {

	_, ok := r.Find(id)

	return ok

}

func (r *Registry) Find(id string) (*Device, bool) {

	for i := range r.Devices {

		if r.Devices[i].ID == id {

			return &r.Devices[i], true

		}

	}

	return nil, false

}

func (r *Registry) Add(device Device) {

	r.Devices = append(
		r.Devices,
		device,
	)

}
