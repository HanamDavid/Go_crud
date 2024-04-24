package task

import "fmt"

type Task struct{
	ID int `json:id`
	Name string `json:"name"`
	Complete bool `json:"complete"`
}


func listTask( tasks []Task){
	if(len(tasks)==0){
		fmt.Println("Empty Task")
		return
	}

	for _,task:= range tasks{
		fmt.Println("%s", task.Name)
	}

}
