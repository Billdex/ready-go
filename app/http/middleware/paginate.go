package middleware

import (
	"github.com/gin-gonic/gin"
	"ready-go/app/http/response"
	"ready-go/pkg/util/numutil"
	"strconv"
)

func Paginate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var page, pageSize int
		var err error
		page, err = strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, err = strconv.Atoi(c.DefaultQuery("page_size", "25"))
		if err != nil {
			response.ErrorValidateParamsFail(c)
			c.Abort()
			return
		}
		page = numutil.IntLeast(page, 1)
		pageSize = numutil.IntBetween(pageSize, 1, 100)
		c.Set("page", page)
		c.Set("page_size", pageSize)
		c.Next()
	}
}
