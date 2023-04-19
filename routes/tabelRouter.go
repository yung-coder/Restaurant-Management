package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func TabelRoutes(incomingRoutes *gin.Engine) { 
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/table/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}