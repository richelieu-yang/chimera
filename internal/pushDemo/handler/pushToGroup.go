package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

// PushToGroup
// @Summary 推送消息给group相关的所有连接（exceptBsids对应的连接例外）.
// @Tag push
// @Accept x-www-form-urlencoded
// @Produce json
// @Param text 			formData	string		true	"推送消息的内容."
// @Param group			formData	string		true	"推送消息的内容."
// @Param exceptBsids	formData	[]string 	false	"例外连接的bsid."
// @Success 200 {object} types.JsonResponse
// @Router /push_to_group [post]
func PushToGroup(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type Params struct {
		Text string `form:"text" binding:"required"`

		Group       string   `form:"group" binding:"required"`
		ExceptBsids []string `form:"exceptBsids" binding:"unique,dive,required"`
	}

	params := &Params{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	if err := pushKit.PushToGroup([]byte(params.Text), params.Group, params.ExceptBsids); err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "no error", nil),
	}, nil
}
