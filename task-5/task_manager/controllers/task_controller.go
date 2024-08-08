package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTasks retrieves all tasks from the data source.
// It returns a JSON response containing the tasks on success,
// or an error message with a status code on failure.
func GetTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a task by its ID.
// It expects the ID to be passed as a parameter in the request URL.
// If the ID is not a valid ObjectID, it returns a 400 Bad Request.
// If the task is not found, it returns a 404 Not Found.
// If any other error occurs, it returns a 500 Internal Server Error.
// The retrieved task is returned as a JSON response with status 200 OK.
func GetTask(c *gin.Context) {

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {  // If the ID is not a valid ObjectID, return a 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}
	task, err := data.GetTaskByID(objID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, task)
}


// CreateTask handles the creation of a new task.
// It expects a JSON payload containing the task details.
// If the payload is valid, it creates a new task and returns the created task as JSON.
// If the payload is invalid or any error occurs during the creation process, it returns an appropriate error message as JSON.
func CreateTask(c *gin.Context) {
	var newTask models.TaskIdLess

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if newTask.Description == ""{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Description can't be empty"})
		return
	}
	if newTask.Title == ""{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Title can't be empty"})
		return
	}
	
	createdTask, err := data.AddNewTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask updates a task with the provided ID.
// It expects a JSON payload containing the updated task information.
// The ID of the task to be updated is extracted from the request parameters.
//
// If the JSON payload cannot be parsed or the ID is not a valid ObjectID, it returns a 400 Bad Request response.
// If the task with the provided ID does not exist, it returns a 404 Not Found response.
// If an error occurs during the update operation, it returns a 500 Internal Server Error response.
// The updated task is returned in the response body if the update is successful.
func UpdateTask(c *gin.Context) {

	var newTask models.TaskIdLess
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{ // If the ID is not a valid ObjectID, return a 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}
	res, err := data.UpdateTaskById(objID, newTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteTask deletes a task by its ID.
//
// Parameters:
// - c: The gin context.
//
// Returns:
// - None.
//
// Behavior:
// - If the ID is not a valid ObjectID, it returns a 400 Bad Request.
// - If the task is not found, it returns a 404 Not Found.
// - If there is an internal server error, it returns a 500 Internal Server Error.
// - If the task is deleted successfully, it returns a 200 OK with a success message.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{ // If the ID is not a valid ObjectID, return a 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}
	_, err = data.DeleteTaskByID(objID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}