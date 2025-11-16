package main

// go package init == nd-KH/taskmanager

import(
	"fmt"
	"os"
	"log"
	tasks "nd-KH/taskmanager/tasks"
)

func main() {
	for{
		fmt.Println("Please choose the option you would like to do:")
		fmt.Print("1. List tasks\n2. Add Tasks\n3. Delete Tasks\n4. Edit Tasks\n\n0. Exit")
		
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			log.Fatal(err)
		}


		makeChoice(int(choice))

	}
}

func makeChoice(x int){
	switch x{
	case 1:
		tasks.ListTasks()
	case 2:
		tasks.AddTasks()
	case 3:
		tasks.DeleteTasks()
	case 4:
		tasks.EditTasks()
	case 0:
		os.Exit(0)
	}
}