package main

import (
	"fmt"
	"net/http"
)

var taskOne = "Wash the dishes"
var taskTwo = "Study for microeconomics exam"
var taskThree = "Watch go crash course"
var taskItems = []string{taskOne, taskTwo, taskThree}

func main() {
	fmt.Println("##### Welcome to our Todolist App! #####")

	// printTasks(taskItems)
	// fmt.Println()

	// taskItems = addTask(taskItems, "Go for a run")

	// fmt.Println()
	// printTasks(taskItems)\

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}
func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}

}
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user, welcome to our to do list app!"
	fmt.Fprintln(writer, greeting)
}
