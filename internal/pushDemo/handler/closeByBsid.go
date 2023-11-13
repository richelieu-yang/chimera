package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

// CloseByBsid
// @Summary 关闭指定连接（根据bsid）.
// @Tags close
// @Accept x-www-form-urlencoded
// @Produce json
// @Param bsid 		formData	string	true	"要关闭哪个连接？"
// @Param reason	formData	string 	false	"关闭的原因."
// @Success 200 {object} types.JsonResponse
// @Router /close_by_bsid [post]
func CloseByBsid(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type closeByBsidParams struct {
		Bsid   string `form:"bsid" binding:"required"`
		Reason string `form:"reason"`
	}

	params := &closeByBsidParams{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	if err := pushKit.CloseByBsid(params.Bsid, params.Reason); err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "no error", nil),
	}, nil
}
