package main
import(
	"github.com/gin-gonic/gin"
	"task_manager/router"
	
) 

func main(){
	r := gin.Default()

	router.SetUpRouter(r)

	r.Run("localhost:8080")
}