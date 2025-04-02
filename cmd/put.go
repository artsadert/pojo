package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/artsadert/pojo/repo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(putTask)
}

var putTask = &cobra.Command{
	// TODO change value of task to smth
	Use:   "put",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Cannot parse arguments")
			os.Exit(1)
		}

		id, err = repo.PutTask(id, args[1])
		if err != nil {
			fmt.Println("Cannot put task", err)
			os.Exit(1)
		}
		fmt.Println("Put task with id:", id)
	},
	Args: cobra.MinimumNArgs(2),
}
