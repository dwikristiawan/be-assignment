package cmd

import (
	"BE-ASSIGMENT/app/config"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	EnvFile string
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "be-assigmnet",
	}
)
var (
	rootConfig *config.Root
)

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&EnvFile, "env", "e", ".env", ".env file to read from")
	fmt.Println(EnvFile)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Cannot Run CLI. err > ", err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(func() {
		configReader()
	})
}
func configReader() {
	log.Infof("Initialize env")
	rootConfig = config.Load(EnvFile)
}
