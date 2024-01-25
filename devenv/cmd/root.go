package cmd

import (
	"embed"
	"fmt"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/elewis787/boa"
	"github.com/metafeather/tools/devenv/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	debug    bool
	global   bool
	cfgFile  string
	cacheDir string
	StdlibFS embed.FS
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devenv",
	Short: "Manage development environment(s)",
	Long:  `Manage your development environment, tools and languages in a reproducible way.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Set debug logging for all subcommands
		if debug {
			log.SetDebug()
		}
		// Log all flags and config settings
		for key, value := range viper.GetViper().AllSettings() {
			log.Debug("settings:", key, value)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	versionTemplate := `{{ printf "%s: %s - version %s\n" .Name .Short .Version }}`
	rootCmd.SetVersionTemplate(versionTemplate)

	// Cobra will make suggestions for mistyped commands and flags
	// rootCmd.DisableSuggestions = false

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "display debug output. (default: false)")
	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().BoolVarP(&global, "global", "g", false, "run global tools (default: false)")
	_ = viper.BindPFlag("global", rootCmd.PersistentFlags().Lookup("global"))
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", path.Join(xdg.ConfigHome, "devenv", "config.yaml"), "config file (default: $XDG_CONFIG_HOME/devenv/config.yaml)")
	// rootCmd.MarkPersistentFlagRequired("config")
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	rootCmd.PersistentFlags().StringVarP(&cacheDir, "cache-dir", "l", path.Join(xdg.CacheHome, "devenv"), "cache dir (default: $XDG_CACHE_HOME/devenv)")
	_ = viper.BindPFlag("cache-dir", rootCmd.PersistentFlags().Lookup("cache-dir"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Bubbletea TUI for help and usage info
	rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.SetHelpFunc(boa.HelpFunc)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".devenv" (without extension).
		viper.SetConfigType("yaml")
		viper.SetConfigName(".devenv")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
	}

	viper.SetEnvPrefix("devenv") // will be uppercased automatically
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// SetVersionInfo sets the version using build or goreleaser metadata
// ref: https://www.jvt.me/posts/2023/05/27/go-cobra-version/
func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}
