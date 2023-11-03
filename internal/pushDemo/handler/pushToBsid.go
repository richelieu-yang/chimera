package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

// PushToBsid
// @Summary 推送消息给指定连接.
// @Tags push
// @Accept x-www-form-urlencoded
// @Produce json
// @Param text	formData	string	true	"推送消息的内容."
// @Param bsid	formData	string 	true	"推送的目标连接."
// @Success 200 {object} types.JsonResponse
// @Router /push_to_bsid [post]
func PushToBsid(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type pushToBsidParams struct {
		Text string `form:"text" binding:"required"`

		Bsid string `form:"bsid" binding:"required"`
	}

	params := &pushToBsidParams{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	if err := pushKit.PushToBsid([]byte(params.Text), params.Bsid); err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "no error", nil),
	}, nil
}
