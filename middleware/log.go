package middleware

import (
	"Memo/conf"
	"Memo/public"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InjectLogID() func(c *gin.Context) {
	return injectLogID
}

func injectLogID(c *gin.Context) {
	// inject log id for tracing
	logID := public.GetUniqueID()
	public.SetLogIDToContext(c, strconv.FormatUint(logID, conf.Decimal))
	public.SetLogID2ResponseHeader(c)
}
