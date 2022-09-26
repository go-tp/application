package model

import (
	"github.com/gin-gonic/gin"
)

type index struct {
	Id int `form:"id" json:"id"`
}

func IndexM(c *gin.Context) interface{} {
	r := index{}
	return r
}