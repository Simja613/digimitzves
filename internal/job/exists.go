package job

import "os"

func Exists() bool {

	_, err := os.Stat(JobPath)

	return err == nil
}
