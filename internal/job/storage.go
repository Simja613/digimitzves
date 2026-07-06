package job

import (
	"encoding/json"
	"os"

	"github.com/digimitzves/core/internal/model"
)

func Save(job model.Job) error {

	data, err := json.MarshalIndent(job, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(JobPath, data, 0644)
}

func Load() (model.Job, error) {

	var job model.Job

	data, err := os.ReadFile(JobPath)
	if err != nil {
		return job, err
	}

	err = json.Unmarshal(data, &job)

	return job, err
}
