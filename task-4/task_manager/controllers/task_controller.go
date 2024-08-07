package controllers

import (
	"task_manager/data"
    "task_manager/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetTasks retrieves all tasks from the data store and returns them as a JSON response.
// It uses the GetAllTasks function from the data package to fetch the tasks.
// The tasks are then serialized into JSON format and sent as the response body.
// The HTTP status code 200 (OK) is set for a successful response.
func GetTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a task by its ID.
// It takes a gin.Context object as a parameter.
// The ID of the task is extracted from the URL parameters of the context.
// If the task is found, it returns the task as a JSON response with status code 200.
// If the task is not found, it returns a JSON response with status code 404 and a message indicating that the task was not found.
func GetTask(c *gin.Context){
	id := c.Param("id")
	task := data.GetTaskByID(id)
	if task == nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return 
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask creates a new task.
// This function is responsible for handling the creation of a new task. It takes a gin.Context object as a parameter.
// The function first binds the JSON data from the request body to a newTask variable using c.ShouldBindJSON method.
// If there is an error in binding the JSON data, it returns a JSON response with the error message.
// Otherwise, it calls the AddNewTask function from the data package to add the new task to the database.
// Finally, it returns a JSON response with the created task.
func CreateTask(c * gin.Context){
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask) ; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	} 
	
	created := data.AddNewTask(newTask)
	c.JSON(http.StatusCreated, created)
}


// UpdateTask updates a task with the specified ID.
// Parameters:
// - c: The gin.Context object for handling HTTP requests and responses.
// Returns:
// This function does not return any values.
// Example:
//   UpdateTask(c *gin.Context)
func UpdateTask(c *gin.Context){

	id := c.Param("id")
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask) ; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	} 
	res := data.UpdateTaskById(id, newTask)
	if res == nil{
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return 
	}
	c.JSON(http.StatusOK, res)
}


// DeleteTask deletes a task by its ID.
//
// Parameters:
// - c: The gin.Context object for handling HTTP requests and responses.
//
// Returns:
// This function does not return any values.
func DeleteTask(c *gin.Context){
	id := c.Param("id")
	res := data.DeleteTaskByID(id)
	if !res{
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
}