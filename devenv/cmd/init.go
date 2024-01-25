package cmd

import (
	"path"

	"github.com/metafeather/tools/devenv/internal"
	"github.com/metafeather/tools/devenv/internal/log"
	"github.com/metafeather/tools/devenv/internal/task"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup this tool",
	Long:  `Setup this tool with default configs and stdlib tools.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := viper.GetString("cache-dir")
		log.Debug("cache-dir:", dir)

		stdlib := internal.NewStdlib(StdlibFS)
		err := stdlib.Write(dir)
		cobra.CheckErr(err)
		log.Debug("stdlib written")

		// Prompt to check tools installed
		cacheDir := viper.GetString("cache-dir")
		args = append([]string{"--taskfile", path.Join(cacheDir, "stdlib", "Taskfile.yaml"), "global:init"}, args...)
		log.Debug("init called:", args)
		return task.Run(args...)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
