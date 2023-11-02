package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

// PushToAll
// @Summary 推送消息给所有连接.
// @Description 推送消息给所有连接（exceptBsids对应的链接例外）.
// @Router /push_to_all [post]
// @Accept application/x-www-form-urlencoded
// @Param text 			formData	string		true	"推送消息的内容."
// @Param exceptBsids	formData	[]string 	false	"例外连接的bsid."
// @Produce json
func PushToAll(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	//if err := httpKit.MakeRequestBodySeekable(ctx.Request); err != nil {
	//	return nil, err
	//}
	//data, err := ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	return nil, err
	//}
	//str := string(data)
	//fmt.Println(str)
	//if err := httpKit.ResetRequestBody(ctx.Request); err != nil {
	//	return nil, err
	//}

	type Params struct {
		Text        string   `form:"text" binding:"required"`
		ExceptBsids []string `form:"exceptBsids,optional"`
	}
	params := &Params{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	if err := pushKit.PushToAll([]byte(params.Text), params.ExceptBsids); err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "ok", nil),
	}, nil
}
