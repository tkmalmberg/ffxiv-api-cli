package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ffxiv",
	Short: "CLI for Final Fantasy XIV",
	Long: `CLI to get information from Final Fantasy XIV.
			Currently set up to return information about specific characters`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName(".config")
	viper.SetConfigType("json")
	viper.SetEnvPrefix("ffxiv")

	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}
