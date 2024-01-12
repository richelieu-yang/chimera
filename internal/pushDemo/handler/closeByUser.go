package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/serialize/json/jsonRespKit"
	"net/http"
)

// CloseByUser
// @Summary 关闭指定连接（根据user）.
// @Tags close
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user 		formData	string	true	"要关闭哪些连接？"
// @Param reason	formData	string 	false	"关闭的原因."
// @Success 200 {object} types.JsonResponse
// @Router /close_by_user [post]
func CloseByUser(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type closeByUserParams struct {
		User   string `form:"user" binding:"required"`
		Reason string `form:"reason"`
	}

	params := &closeByUserParams{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	pushKit.CloseByUser(params.User, params.Reason)
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "no error", nil),
	}, nil
}
