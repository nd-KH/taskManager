package logic

import(
	"fmt"
	"os"
	"encoding/json"
)

type Task struct {
	name			string
	description 	string
	progress		int
	complete		bool
}

func LoadJson()([]Task, error) {
	data, err := os.ReadFile("taskData.json")
	if err != nil {
		return nil, err
	}
	var Tasks []Task
	if err := json.Unmarshal(data, &Tasks); err != nil {
		return nil, err
	}
	return Tasks, nil
}

func SaveJson(Tasks []Task)(error){
	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("taskData.json", data, 0644)
}

func CreateFile()error {
	file, err := os.Create("taskData.json")
	fmt.Println("Wrote file: ", file)
	if err != nil {
		return err
	}
	return nil
}
