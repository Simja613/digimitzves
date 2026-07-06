package model

type JobStatus string

const (
	JobPlanned JobStatus = "planned"

	JobReady JobStatus = "ready"

	JobRunning JobStatus = "running"

	JobCompleted JobStatus = "completed"

	JobFailed JobStatus = "failed"

	JobInvalid JobStatus = "invalid"
)

type CommandStatus string

const (
	CommandPlanned CommandStatus = "planned"

	CommandRunning CommandStatus = "running"

	CommandDone CommandStatus = "done"

	CommandFailed CommandStatus = "failed"
)
