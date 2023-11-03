package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

func PushToUser(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type Params struct {
		Text string `form:"text" binding:"required"`

		User        string   `form:"user" binding:"required"`
		ExceptBsids []string `form:"exceptBsids" binding:"unique,dive,required"`
	}

	params := &Params{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	if err := pushKit.PushToUser([]byte(params.Text), params.User, params.ExceptBsids); err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "ok", nil),
	}, nil
}
