package task_test

import (
	"testing"

	"github.com/metafeather/tools/devenv/internal/task"
)

func TestRun(t *testing.T) {
	_ = task.Run([]string{}...)
}
