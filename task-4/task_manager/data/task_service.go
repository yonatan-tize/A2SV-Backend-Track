package data

import (
	"task_manager/models"
)

var tasks = []models.Task{
    {
        ID:          "1",
        Title:       "Buy Groceries",
        Description: "Milk, Bread, Cheese, Eggs",
        DueDate:     "2024-08-10",
        Status:      "Pending",
    },
    {
        ID:          "2",
        Title:       "Complete Homework",
        Description: "Math, Science, History",
        DueDate:     "2024-08-09",
        Status:      "Pending",
    },
    {
        ID:          "3",
        Title:       "Doctor Appointment",
        Description: "Annual check-up",
        DueDate:     "2024-08-12",
        Status:      "Pending",
    },
}

// GetAllTasks returns all the tasks in the task manager.
func GetAllTasks() []models.Task{
	return tasks
}
// function to find a specific task using the id
// GetTaskByID retrieves a task from the task list based on the provided ID.
// It searches for a task with a matching ID and returns a pointer to the task if found.
// If no task is found with the given ID, it returns nil.

func GetTaskByID(id string) *models.Task{

	for _, task := range tasks{
		if id == task.ID{
			return &task
		}
	}
	return nil
}


// adding a new task 
// AddNewTask adds a new task to the task list.
// It takes a task of type models.Task as a parameter and returns a pointer to the added task.
func AddNewTask(task models.Task) *models.Task{
	tasks = append(tasks, task)
	return &task
}


// UpdateTaskById updates a task in the task list with the specified ID.
// It takes the ID of the task to be updated and the updatedTask as input parameters.
// If a task with the specified ID is found, it replaces the existing task with the updatedTask and returns a pointer to the updated task.
// If no task with the specified ID is found, it returns nil.
func UpdateTaskById(id string, updatedTask models.Task) *models.Task {

	for i, task := range tasks{
		if task.ID == id{
			tasks[i] = updatedTask
			return &tasks[i]
		}
	}
	return nil

}

// DeleteTaskByID deletes a task from the task list based on the provided ID.
// It searches for the task with the matching ID and removes it from the tasks slice.
// If a task is successfully deleted, it returns true. Otherwise, it returns false.
func DeleteTaskByID(id string) bool{
	for i, task := range tasks{
		if task.ID == id{
			tasks =append(tasks[:i], tasks[i+1:]...) 
			return true
		}
	}
	return false
}