package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "tes",
}

var childA = &cobra.Command{
	Use: "a",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("config flag cmd a: '%s'", viper.GetString("config"))
	}}
var childB = &cobra.Command{
	Use: "b",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("config flag cmd b: '%s'", viper.GetString("config"))
	},
}

func init() {
	rootCmd.AddCommand(childA)
	rootCmd.AddCommand(childB)

	ConfigurationString(childA, "config", "c", "", "Path to configuration file")
	ConfigurationString(childB, "config", "c", "", "Path to configuration file")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func ConfigurationString(command *cobra.Command, key string, shorthand string, defaultValue string, usage string) {
	command.Flags().StringP(key, shorthand, defaultValue, usage)
	_ = viper.BindPFlag(key, command.Flags().Lookup(key))
}
