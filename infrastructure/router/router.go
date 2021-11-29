package routes

import (
	"github.com/Le-MaliX/ACADEMY-GO-Q42021/interface/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.Home)

	r.GET("/monsters", controllers.GetAllMonsters)

	r.GET("/monsters/:id", controllers.GetMonsterById)

	return r
}
