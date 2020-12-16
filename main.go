package main

import (
	"chatroom/routers/webRouter"
	"chatroom/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {

	service := gin.Default()
	webRouter.Register(service)
	web_port := utils.GetVal("web","port")
	http.ListenAndServe(":"+web_port, service)

}