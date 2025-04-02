package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/artsadert/pojo/config"
	"github.com/artsadert/pojo/models"
)

func CreateTaskJSON(value string) (uint64, error) {
	id, err := get_last_id()
	if err != nil {
		return 0, errors.New(fmt.Sprint("Cannot get id of last task\n", err))
	}

	new_task := models.Task{ID: id, Value: value, Completed: false}
	tasks, err := GetTasks()
	if err != nil {
		return 0, errors.New(fmt.Sprint("Cannot get all tasks"))
	}

	tasks = append(tasks, new_task)

	data, err := json.Marshal(tasks)
	if err != nil {
		return 0, errors.New(fmt.Sprint("Cannot serialize model\n", err))
	}

	os.WriteFile(config.StoringPath, data, 0666)

	return id, nil
}
