package router

import(
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
	"task_manager/middleware"
)

func SetUpRouter(router *gin.Engine){

	router.POST("/register", controllers.CreateAccount)
	router.POST("/login", controllers.Login)


	// we have to create the middleware
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/tasks", controllers.GetTasks)
	auth.GET("/tasks/:id", controllers.GetTask)

	auth = router.Group("/admin")
	
	auth.Use(middleware.AuthAdminMiddleware())
	auth.POST("/promote/:id", controllers.PromoteUser)
	auth.POST("/tasks", controllers.CreateTask)
	auth.PUT("/tasks/:id", controllers.UpdateTask)
	auth.DELETE("/tasks/:id", controllers.DeleteTask)
	
}