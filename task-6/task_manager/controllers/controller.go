package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateAccount creates a new user account.
//
// This function receives a JSON payload containing the user information and binds it to the newUser variable.
// If the binding fails, it returns a JSON response with a bad request error.
// Otherwise, it calls the CreateAccount function from the data package to create the user account.
// If an error occurs during the account creation, it returns a JSON response with the error message.
// Finally, it returns a JSON response with the created user account if the account creation is successful.
func CreateAccount(c *gin.Context){

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}
	user, err := data.CreateAccount(newUser)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, user)
}

// Login handles the login functionality for the task manager.
// It expects a JSON payload containing the username and password.
// If the payload is valid, it authenticates the user and returns a token and user information.
// If the authentication fails, it returns an error message.
// The token and user information are returned in a JSON response with status code 202 (Accepted).
// If there is an error in binding the JSON payload, it returns an error message with status code 400 (Bad Request).
func Login(c *gin.Context){
	// by using username and password
	var user models.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token,  err := data.AuthenticateUser(user.Username, user.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
	}
	c.JSON(http.StatusAccepted, gin.H{"token": token, "user": user})
}

// PromoteUser promotes a user to an admin role.
// It takes a gin.Context as a parameter and retrieves the user ID from the request parameters.
// If the ID is not a valid ObjectID, it returns a 400 Bad Request response.
// It then calls the UpdateUserRoll function to update the user's role to admin.
// If the user is not found, it returns a 404 Not Found response.
// If there is any other error during the update process, it returns a 400 Bad Request response.
// Finally, it returns a 200 OK response with a message indicating that the user has been promoted to admin.
func PromoteUser(c *gin.Context){
	userId := c.Param("id") 

	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil{ // If the ID is not a valid ObjectID, return a 400 Bad Request
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}
	
	err = data.UpdateUserRoll(objID)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return 
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "promoted to admin"})
}


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
	var newTask models.Task

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

	var newTask models.Task
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