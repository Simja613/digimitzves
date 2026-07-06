package health

type Status int

const (
	OK Status = iota
	WARNING
	ERROR
)

type Service struct {
	timeStatus     Status
	scheduleStatus Status
	deviceStatus   Status
}

func New() *Service {
	return &Service{
		timeStatus:     OK,
		scheduleStatus: OK,
		deviceStatus:   OK,
	}
}

func (s *Service) Overall() Status {

	if s.timeStatus == ERROR ||
		s.scheduleStatus == ERROR ||
		s.deviceStatus == ERROR {
		return ERROR
	}

	if s.timeStatus == WARNING ||
		s.scheduleStatus == WARNING ||
		s.deviceStatus == WARNING {
		return WARNING
	}

	return OK
}
