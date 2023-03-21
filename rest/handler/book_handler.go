package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yonisaka/book-service/consts"
	"github.com/yonisaka/book-service/rest/dto"
	pbAuth "github.com/yonisaka/protobank/auth"
	"net/http"
)

type BookHandler struct {
	*Handler
}

type BookRequest struct {
}

func NewBookHandler(h *Handler) *BookHandler {
	return &BookHandler{h}
}

func (r *BookHandler) GetBookList(c *gin.Context) {
	authb2b := pbAuth.AuthB2BPayload{
		Token: c.Request.Header.Get("Authorization"),
	}
	_, err := r.client.AuthB2B(c, &authb2b)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Code:    http.StatusUnauthorized,
			Message: fmt.Sprintf("%v", consts.MessageUnauthorized),
		})
		return
	}

	rsp, err := r.repo.Book.Get(c)
	//rsp, err := r.client.GetBookList(c, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: consts.MessageSuccess,
		Data:    rsp,
	})
	return
}
