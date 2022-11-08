package refererKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewGinRefererMiddleware(builders []*RefererVerifierBuilder) (gin.HandlerFunc, error) {
	verifiers := make([]*RefererVerifier, 0, len(builders))
	for _, builder := range builders {
		if builder != nil {
			verifier, err := builder.Build()
			if err != nil {
				return nil, err
			}
			verifiers = append(verifiers, verifier)
		}
	}

	middleware := func(ctx *gin.Context) {
		for _, verifier := range verifiers {
			if ok, reason := verifier.VerifyByGinContext(ctx); !ok {
				// 验证失败
				// http.StatusUnauthorized: 401
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"reason": reason,
				})
				return
			}
		}
		// 验证成功
		ctx.Next()
	}
	return middleware, nil
}
