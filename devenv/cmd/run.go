package cmd

import (
	"path"

	"github.com/metafeather/tools/devenv/internal/log"
	"github.com/metafeather/tools/devenv/internal/task"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO(metafeather): eval https://github.com/goyek/goyek
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a task",
	Long:  `Run tasks defined in a local Taskfile.yaml (see https://taskfile.dev)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		debug := viper.GetBool("debug")
		if debug {
			// Override default taskfile --silent flag
			args = append([]string{"--silent=false"}, args...)
		}

		// Use global tasks from stdlib
		global := viper.GetBool("global")
		cacheDir := viper.GetString("cache-dir")
		if global {
			args = append([]string{"--taskfile", path.Join(cacheDir, "stdlib", "Taskfile.yaml")}, args...)
		}

		log.Debug("run called:", args)
		return task.Run(args...)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
