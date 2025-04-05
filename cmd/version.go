package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	// creates task
	Use:   "version",
	Short: "Gets version of cli",
	Long:  "Gets version of cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pojo 0.1")
	},
}

func init() {
	rootCmd.AddCommand(Version)
}
