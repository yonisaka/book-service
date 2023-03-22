package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yonisaka/book-service/consts"
	"github.com/yonisaka/book-service/domain/entity"
	"github.com/yonisaka/book-service/rest/dto"
	"net/http"
	"strconv"
)

type BookHandler struct {
	*Handler
}

func NewBookHandler(h *Handler) *BookHandler {
	return &BookHandler{h}
}

func (r *BookHandler) GetBookList(c *gin.Context) {
	rsp, err := r.repo.Book.Get(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		*dto.NewResponse().WithCode(http.StatusOK).WithData(rsp),
	)
	return
}

func (r *BookHandler) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rsp, err := r.repo.Book.Find(c, id)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		*dto.NewResponse().WithCode(http.StatusOK).WithData(rsp),
	)
	return
}

func (r *BookHandler) CreateBook(c *gin.Context) {
	var req entity.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			*dto.NewResponse().WithCode(http.StatusBadRequest).WithMessage(err.Error()),
		)
		return
	}
	req.Author = c.GetString("username")
	err := r.repo.Book.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		*dto.NewResponse().WithCode(http.StatusOK).WithMessage(consts.MessageSuccessCreate).WithData(req),
	)
	return
}

func (r *BookHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req entity.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			*dto.NewResponse().WithCode(http.StatusBadRequest).WithMessage(err.Error()),
		)
		return
	}
	req.Author = c.GetString("username")
	err := r.repo.Book.Update(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage(err.Error()),
		)
		return
	}

	req.ID = uint(id)
	c.JSON(
		http.StatusOK,
		*dto.NewResponse().WithCode(http.StatusOK).WithMessage(consts.MessageSuccessUpdate).WithData(req),
	)
	return
}

func (r *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := r.repo.Book.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		*dto.NewResponse().WithCode(http.StatusOK).WithMessage(consts.MessageSuccessDelete),
	)
	return
}
