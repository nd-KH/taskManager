package logic

import(
	"fmt"
)

type Task struct {
	name			string
	description 	string
	progress		int
	complete		bool
}

func LoadJson()([]Task, err) {
	data, err := os.Readfile("taskData.json")
	if err != nil {
		return nil, err
	}
	var Tasks []Task
	if err := json.Unmarshal(data, &Tasks); err != nil {
		return nil, err
	}
	return Tasks, nil
}

func SaveJson(Tasks []Task)(err){
	data, err := json.Marshalindent(Tasks, "", "  ")
	if err != nil {
		return nil, err
	}
	return os.Writefile("taskData.json", data, 0644)
}

func CreateFile()(err) {
	file, err := os.Create("taskData.json")
	if err != nil {
		return err
	}
	return nil
}
