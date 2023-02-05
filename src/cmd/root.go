package cmd

import (
	util "cli/src/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var BuildConfig util.Config
var RunTimeConfig string

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "awesome-cli",
	Short: "Awesome CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		RunTimeConfig = util.GetRuntimeConfig(viper.ConfigFileUsed())
		viper.Unmarshal(&BuildConfig)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Awesome CLI v1.0.0")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/awesome-cli.yml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("awesome-cli")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("$PWD")
	}

	viper.ReadInConfig()

	data := util.GetBuildConfig(viper.ConfigFileUsed())
	util.ReadInConfig(data)
	util.ReplaceEnvVars()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
