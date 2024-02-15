package middleware

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(ac appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		context.Next()
	}
}
