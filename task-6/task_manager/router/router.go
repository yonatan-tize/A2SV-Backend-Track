package router

import(
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
	"task_manager/middleware"
)

func SetUpRouter(router *gin.Engine){

	router.POST("/register", controllers.CreateAccount)
	router.POST("/login", controllers.Login)

	router.GET("/tasks", middleware.AuthMiddleware(), controllers.GetTasks)
	router.GET("/tasks/:id", middleware.AuthMiddleware(),  controllers.GetTask)

	router.PUT("/admin/promote/:id", middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(), controllers.PromoteUser)
	router.POST("/admin/tasks", middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(), controllers.CreateTask)
	router.PUT("/admin/tasks/:id", middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(), controllers.UpdateTask)
	router.DELETE("/admin/tasks/:id", middleware.AuthMiddleware(), middleware.AuthAdminMiddleware(), controllers.DeleteTask)
	
}