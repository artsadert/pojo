package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/artsadert/pojo/config"
	"github.com/artsadert/pojo/repo"
)

func init() {
	rootCmd.AddCommand(createTask)
}

var createTask = &cobra.Command{
	// creates task
	Use:   "create",
	Short: "Creates tasks",
	Long:  "Creates tasks and put it to config file " + config.StoringPath,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			id, err := repo.CreateTaskJSON(arg)
			if err != nil {
				fmt.Printf("Error task JSON\n%s", err)
				os.Exit(1)
			}
			fmt.Printf("Created task with id %d\n", id)
		}
	},
	Args: cobra.MinimumNArgs(1),
}
