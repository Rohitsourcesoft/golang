package main
import "github.com/gin-gonic/gin"
import "github.com/golang/controllers"
func main() {
   router := gin.Default()
   router.GET("/", func(c *gin.Context) {
       c.JSON(200, gin.H{
           "message": "Welcome to Go lang!!",
       })
   })
   router.GET("/login", controllers.LoginUser)



   
   router.Run(":9999")
}
