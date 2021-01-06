package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/jstang9527/buoy/utils"
)

// TrojanItemOutput 木马服务
type TrojanItemOutput struct {
	AssetID   string `json:"asset_id" form:"asset_id"`
	AssetName string `json:"asset_name" form:"asset_name"`
	IsConn    int8   `json:"is_Conn" form:"is_Conn"`
}

// TrojanListOutput ...
type TrojanListOutput struct {
	Total int64               `json:"total" form:"total" comment:"总数" example:"" validate:""` //总数
	List  []*TrojanItemOutput `json:"list" form:"list" comment:"列表" example:"" validate:""`   //列表
}

// TrojanConnInput 木马连接会话
type TrojanConnInput struct {
	AssetID   string `json:"id" form:"id" comment:"资产ID" example:"asset_id(\\w{24})" validate:"required"`
	AssetName string `json:"name" form:"name" comment:"资产名" example:"asset_name(weblogic)" validate:"required"`
	SpareLine int8   `json:"line" form:"line" comment:"连接线路" example:"0" validate:"required"`
}

// BindValidParam 校验新增参数,绑定结构体,校验参数
func (s *TrojanConnInput) BindValidParam(c *gin.Context) error {
	return utils.DefaultGetValidParams(c, s)
}
