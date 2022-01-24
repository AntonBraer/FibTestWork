package api

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func NewError(ctx *gin.Context, status int, err error) {
	er := errorResponse{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}
