package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jstang9527/buoy/middleware"
)

// DemoController ...
type DemoController struct{}

// DemoRegister ...
func DemoRegister(group *gin.RouterGroup) {
	demo := &DemoController{}
	group.GET("/info", demo.DemoInfo)
}

// DemoInfo godoc
// @Summary 模板-Demo
// @Description 模板简介
// @Tags 模板
// @ID /demo/info
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /demo/info [get]
func (d *DemoController) DemoInfo(c *gin.Context) {
	out := "demoinfo"
	middleware.ResponseSuccess(c, out)
}
