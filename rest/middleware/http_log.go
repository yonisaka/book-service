package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonisaka/book-service/grpc/client"
	"github.com/yonisaka/book-service/rest/dto"
	pbLog "github.com/yonisaka/protobank/log"
)

func SaveHttpLog(client *client.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := pbLog.SaveHttpLogRequest{
			Ip:     c.ClientIP(),
			Path:   c.Request.RequestURI,
			Method: c.Request.Method,
		}

		_, err := client.SaveHttpLog(c, &payload)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				*dto.NewResponse().WithCode(http.StatusUnprocessableEntity).WithMessage(err.Error()),
			)
			return
		}
		c.Next()
	}
}
