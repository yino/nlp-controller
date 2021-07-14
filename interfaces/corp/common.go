package corp

import "github.com/gin-gonic/gin"

func GetUid(c *gin.Context) uint64 {
	return c.GetUint64("uid")
}
