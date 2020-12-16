package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Index struct {

}

func (i *Index)Index(c *gin.Context){
	c.JSON(http.StatusOK, 1111)

}