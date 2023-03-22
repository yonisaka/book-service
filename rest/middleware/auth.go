package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yonisaka/book-service/grpc/client"
	"github.com/yonisaka/book-service/rest/dto"
	pbAuth "github.com/yonisaka/protobank/auth"
	"net/http"
)

func AuthB2B(client *client.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := pbAuth.AuthB2BPayload{
			Token: c.Request.Header.Get("Authorization"),
		}
		user, err := client.AuthB2B(c, &payload)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				*dto.NewResponse().WithCode(http.StatusUnauthorized).WithMessage(err.Error()),
			)
			return
		}
		c.Set("username", user.Username)
		c.Next()
	}
}
