package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

// CloseAll
// @Summary 关闭所有连接.
// @Tags close
// @Accept x-www-form-urlencoded
// @Produce json
// @Param reason formData string false "关闭的原因."
// @Success 200 {object} types.JsonResponse
// @Router /close_all [post]
func CloseAll(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type closeAllParams struct {
		Reason string `form:"reason"`
	}

	params := &closeAllParams{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	pushKit.CloseAll(params.Reason)
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "no error", nil),
	}, nil
}
