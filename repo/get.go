package repo

import (
	"encoding/json"
	"os"

	"github.com/artsadert/pojo/config"
	"github.com/artsadert/pojo/models"
)

func GetTasks() ([]models.Task, error) {
	fileBytes, _ := os.ReadFile(config.StoringPath)
	tasks := []models.Task{}

	if err := json.Unmarshal(fileBytes, &tasks); err != nil {
		//return tasks, errors.New(fmt.Sprintf("Error while loading file!\n%s", err))
		return tasks, nil
	}

	return tasks, nil
}
