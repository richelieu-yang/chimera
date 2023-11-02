package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
)

func PushToAll(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type Params struct {
		Text        string   `form:"text"`
		ExceptBsids []string `form:"exceptBsids,optional"`
	}
	params := &Params{}
	if err := ctx.Bind(params); err != nil {
		return nil, err
	}

	err := pushKit.PushToAll([]byte(params.Text), params.ExceptBsids)
	if err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "ok", nil),
	}, nil
}
