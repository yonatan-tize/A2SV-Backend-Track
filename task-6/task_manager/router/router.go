package router

import(
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
	"task_manager/middleware"
)

func SetUpRouter(router *gin.Engine) {
	// Public routes
	public := router.Group("/")
	{
		public.POST("/register", controllers.CreateAccount)
		public.POST("/login", controllers.Login)
	}

	// Authenticated routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/tasks", controllers.GetTasks)
		authorized.GET("/tasks/:id", controllers.GetTask)
	}

	// Admin routes (require admin privileges)
	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AuthAdminMiddleware())
	{
		admin.PUT("/promote/:id", controllers.PromoteUser)
		admin.POST("/tasks", controllers.CreateTask)
		admin.PUT("/tasks/:id", controllers.UpdateTask)
		admin.DELETE("/tasks/:id", controllers.DeleteTask)
	}
}
