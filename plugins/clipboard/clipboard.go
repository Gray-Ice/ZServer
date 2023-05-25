package clipboard

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Clipboard struct {
}

func (c *Clipboard) Name() string {
	return "clipboard"
}

func (c *Clipboard) Timeout() time.Duration {
	return 4 * time.Second
}

func (c *Clipboard) Reset(flag int) {
}
func (c *Clipboard) PhoneRequestHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"text": "This is my clipboard"})
}
