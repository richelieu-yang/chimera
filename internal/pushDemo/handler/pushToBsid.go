package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

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
		Object: jsonRespKit.PackFully("0", "ok", nil),
	}, nil
}
