package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

// CloseById
// @Summary 关闭指定连接（根据id）.
// @Tags close
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id 		formData	string	true	"要关闭哪个连接？"
// @Param reason	formData	string 	false	"关闭的原因."
// @Success 200 {object} types.JsonResponse
// @Router /close_by_id [post]
func CloseById(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type closeByIdParams struct {
		Id     string `form:"id" binding:"required"`
		Reason string `form:"reason"`
	}

	params := &closeByIdParams{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	if err := pushKit.CloseById(params.Id, params.Reason); err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "no error", nil),
	}, nil
}
