package data

import (
	"context"
	"fmt"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var client *mongo.Client
var collection *mongo.Collection

// init initializes the MongoDB client and establishes a connection to the database.
// It sets up the necessary configurations and checks if the connection is successful.
// The MongoDB connection URI is set to "mongodb://localhost:27017".
// If any error occurs during the initialization process, it will be logged as fatal.
// After a successful connection, the "taskManager" database and "tasks" collection are selected.
func init(){
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal("database not connected")
	}
	fmt.Println("connected to database")
	collection = client.Database("taskManager").Collection("tasks")
}

// GetAllTasks retrieves all tasks from the database.
// It returns a slice of models.Task and an error if any.
func GetAllTasks() ([]models.Task, error){
	var tasks []models.Task
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil{
		return nil, err
	}
	defer cursor.Close(context.TODO())
	
	for cursor.Next(context.TODO()){
		var task models.Task
		err := cursor.Decode(&task) 
		if err != nil{
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil{
		return nil, err
	}
	return tasks, nil
}

// GetTaskByID retrieves a task from the database based on the provided ID.
// It takes an `id` parameter of type `primitive.ObjectID` and returns a pointer to a `models.Task` and an error.
// If the task is found, it returns a pointer to the task and a `nil` error.
// If no task is found, it returns `nil` and an error of type `mongo.ErrNoDocuments`.
// If an error occurs during the retrieval process, it returns `nil` and the corresponding error.
func GetTaskByID(id primitive.ObjectID) (*models.Task, error){
	filter := bson.M{"_id": id}

	var task models.Task
    err := collection.FindOne(context.TODO(), filter).Decode(&task)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, err // No document found
        }
        return nil, err
    }
    return &task, nil
}


// AddNewTask adds a new task to the collection.
// It takes a task of type models.TaskIdLess as input and returns a pointer to the inserted task and an error, if any.
// The task is inserted into the collection using the InsertOne method.
// The inserted ID is retrieved and used to create a new task with the same properties as the input task.
// The new task is then returned along with a nil error if the insertion is successful.
// If there is an error during the insertion, nil is returned for the task and the error is returned.
func AddNewTask(task models.TaskIdLess) (*models.Task, error) {
    // Insert the task into the collection
    result, err := collection.InsertOne(context.TODO(), task)
    if err != nil {
        return nil, err
    }
    // Get the inserted ID directly as a primitive.ObjectID
    insertedID := result.InsertedID.(primitive.ObjectID)
    // Create a new task with the inserted ID
    insertedTask := models.Task{
        ID:          insertedID,
        Title:       task.Title,
        Description: task.Description,
        DueDate:     task.DueDate,
        Status:      task.Status,
    }
    return &insertedTask, nil
}


// UpdateTaskById updates a task in the database with the specified ID.
// It takes the ID of the task to be updated and the updatedTask object containing the new values.
// The function returns the updated task and an error, if any.
func UpdateTaskById(id primitive.ObjectID, updatedTask models.TaskIdLess) (*models.Task, error) {
    // Create the update document
    update := bson.M{
        "$set": bson.M{
            "title":       updatedTask.Title,
            "description": updatedTask.Description,
            "due_date":    updatedTask.DueDate,
            "status":      updatedTask.Status,
        },
    }

    filter := bson.M{"_id": id}
    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
    var updated models.Task
    err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updated)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, err
        }
        return nil, err
    }
    return &updated, nil
}

// DeleteTaskByID deletes a task from the collection by its ID.
// It takes the ID of the task as a parameter and returns a boolean value indicating whether the task was deleted successfully or not, along with any error that occurred during the deletion process.
func DeleteTaskByID(id primitive.ObjectID) (bool, error){

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil{
		return false, err
	}
	return true, nil
}
