package router

import(
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

func SetUpRouter()*gin.Engine{
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.POST("/tasks", controllers.CreateTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)

	router.DELETE("/tasks/:id", controllers.DeleteTask)
	
	return router

}