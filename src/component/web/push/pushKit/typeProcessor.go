package pushKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Processor interface {
		HandleWithGin(ctx *gin.Context)

		Handle(w http.ResponseWriter, r *http.Request)
	}
)
