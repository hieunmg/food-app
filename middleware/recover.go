package middleware

import (
	"food_delivery/components/appctx"
	"food_delivery/modules/common"

	"github.com/gin-gonic/gin"
)

func Recover(ctx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func ()  {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appError, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appError.StatusCode, appError)

					panic(err)
				}

				appError := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appError.StatusCode, appError)
				panic(err)
			}
		}()
		
		c.Next()
	}
}