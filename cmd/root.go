package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Pojo is a task_manager app",
	Long:  "Pojo is a task_manager app that provides comfortable commands for",
	Run: func(cmd *cobra.Command, args []string) {
		for _, x := range args {
			fmt.Println(x)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
