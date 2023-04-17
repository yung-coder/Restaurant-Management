package routes

import "github.com/gin-gonic/gin"

func TabelRoutes(incomingRoutes *gin.Engine) { 
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/table/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
}