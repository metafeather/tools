package cmd

import (
	"path"

	"github.com/metafeather/tools/devenv/internal/log"
	"github.com/metafeather/tools/devenv/internal/task"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doctorCmd represents the doctor command
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check for potential problems",
	Long: `Check your devenv for potential problems and
report a summary.
	
Note: these warnings are just used to help with debugging.
If everything is working fine: please don't worry or
file an issue; just ignore this.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cacheDir := viper.GetString("cache-dir")
		taskfile := path.Join(cacheDir, "embed", "stdlib", "Taskfile.yaml")
		args = append([]string{"--taskfile", taskfile, "global:doctor"}, args...)
		log.Debug("doctor called:", args)
		return task.Run(args...)
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
