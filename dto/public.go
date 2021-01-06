package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/jstang9527/buoy/utils"
)

// PublicIDInput 可做删除/查询ID的公共模型
type PublicIDInput struct {
	ID int64 `json:"id" form:"id" comment:"target_id" example:"22" validate:"required"` // 服务ID
}

// BindValidParam 校验新增参数,绑定结构体,校验参数
func (s *PublicIDInput) BindValidParam(c *gin.Context) error {
	return utils.DefaultGetValidParams(c, s)
}

// PublicListInput 可做列表分页查询的公共模型
type PublicListInput struct {
	Info  string `json:"info" form:"info" comment:"关键词" example:"" validate:""`              //关键词
	Page  int    `json:"page" form:"page" comment:"页数" example:"1" validate:"required"`      //页数
	Limit int    `json:"limit" form:"limit" comment:"每页条数" example:"20" validate:"required"` // 每页条数
}

// BindValidParam 校验新增参数,绑定结构体,校验参数
func (s *PublicListInput) BindValidParam(c *gin.Context) error {
	return utils.DefaultGetValidParams(c, s)
}