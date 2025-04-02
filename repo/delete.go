package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/artsadert/pojo/config"
)

func DeleteTask(id uint64) (uint64, error) {
	tasks, err := GetTasks()
	if err != nil {
		return 0, errors.New(fmt.Sprint("Cannot get all tasks"))
	}
	found := false

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return 0, errors.New(fmt.Sprintf("Cannot find task with id %d\n", id))
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		return 0, errors.New(fmt.Sprint("Cannot serialize model\n", err))
	}

	os.WriteFile(config.StoringPath, data, 0666)

	return id, nil

}
