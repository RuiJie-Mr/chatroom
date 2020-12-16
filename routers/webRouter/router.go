package webRouter

import (
	"chatroom/controller"
	"github.com/gin-gonic/gin"

)
func Register(service *gin.Engine){
	service.LoadHTMLGlob("templates/*")
	service.GET("/index",new(controller.Index).Index)
}