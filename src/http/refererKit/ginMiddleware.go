package refererKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewGinRefererMiddlewares
/*
@return err == nil的情况下，第一个返回值必定不为nil
*/
func NewGinRefererMiddlewares(builders []*RefererVerifierBuilder) ([]gin.HandlerFunc, error) {
	middlewares := make([]gin.HandlerFunc, 0, len(builders))

	for _, builder := range builders {
		if builder != nil {
			verifier, err := builder.Build()
			if err != nil {
				return nil, err
			}
			middlewares = append(middlewares, func(ctx *gin.Context) {
				if ok, reason := verifier.VerifyByGinContext(ctx); ok {
					ctx.Next()
				} else {
					// http.StatusUnauthorized: 401
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"reason": reason,
					})
				}
			})
		}
	}
	return middlewares, nil
}
