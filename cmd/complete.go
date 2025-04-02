package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/artsadert/pojo/repo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(CompleteTask)
}

var CompleteTask = &cobra.Command{
	// marks task as compeled
	Use:   "complete",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Cannot parse arguments")
			os.Exit(1)
		}
		id, err = repo.CompleteTask(id)
		if err != nil {
			fmt.Println("Cannot complete task", err)
			os.Exit(1)
		}
		fmt.Println("Completed task with id:", id)
	},
	Args: cobra.MinimumNArgs(1),
}
