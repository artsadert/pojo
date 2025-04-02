package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/artsadert/pojo/config"
	"github.com/artsadert/pojo/models"
)

func get_last_id() (uint64, error) {
	fileBytes, err := os.ReadFile(config.StoringPath)
	if os.IsNotExist(err) {
		return 0, nil
	}
	tasks := []models.Task{}

	if err := json.Unmarshal(fileBytes, &tasks); err != nil {
		return 0, errors.New(fmt.Sprintln("Error while loading file!", err))
	}

	var res uint64 = 0
	for _, task := range tasks {
		res = task.ID
	}
	return res + 1, nil
}
