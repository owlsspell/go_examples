package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/owlsspell/todo-app"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON((&input)); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signIn(c *gin.Context) {
}
