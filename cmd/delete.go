package cmd

import (
	"fmt"
	"github.com/artsadert/pojo/repo"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(deleteTask)
}

var deleteTask = &cobra.Command{
	// deletes task from task list
	Use:   "delete",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Cannot parse arguments")
			os.Exit(1)
		}
		id, err = repo.DeleteTask(id)
		if err != nil {
			fmt.Println("Cannot delete task", err)
			os.Exit(1)
		}
		fmt.Println("Deleted task with id:", id)
	},
	Args: cobra.MinimumNArgs(1),
}
