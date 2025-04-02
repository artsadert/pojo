package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(Version)
}

var Version = &cobra.Command{
	// creates task
	Use:   "version",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pojo 0.1")
	},
}
