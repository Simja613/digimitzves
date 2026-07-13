package engine

type Context struct {
	RecoveryRequired bool

	SynchronizationRequired bool

	ExecutionRequired bool

	SaveRequired bool

	DetectedMode SystemMode

	NewDevices bool

	MissingDevices bool

	ConfigurationRequired bool
}
