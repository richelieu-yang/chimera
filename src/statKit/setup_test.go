package statKit

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMustSetup(t *testing.T) {
	MustSetup("_stat.log")

	engine := gin.Default()
	if err := engine.Run(":8888"); err != nil {
		panic(engine)
	}
}
