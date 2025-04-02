package cmd

import (
	"fmt"
	"os"

	"github.com/artsadert/pojo/repo"
	"github.com/spf13/cobra"
)

var a_option bool

var getTask = &cobra.Command{
	// TODO change value of task to smth
	Use:   "get",
	Short: "Gets tasks",
	Long:  "Gets tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := repo.GetTasks()
		if err != nil {
			fmt.Println("An error occupied during reading file", err)
			os.Exit(1)
		}
		for _, task := range tasks {
			fmt.Printf("\033[0m%-5d", task.ID)
			if !task.Completed {
				fmt.Printf("\033[1;32m%s\n", task.Value)
			} else if a_option {
				fmt.Printf("\033[1;31m%s\n", task.Value)
			}
		}
	},
}

func init() {
	getTask.Flags().BoolVarP(&a_option, "all", "a", false, "To get all tasks (already completed)")
	rootCmd.AddCommand(getTask)
}
