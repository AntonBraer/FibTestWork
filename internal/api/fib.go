package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getFibRequest struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type getFibResponse struct {
	Sequence string
}

// @Summary      getFib
// @Tags         fib
// @ID           get-fib
// @Description  get Fibonacci sequence
// @Accept       json
// @Produce      json
// @Param        input  body      getFibRequest  true  "start and end"
// @Success      200    {object}  getFibResponse
// @Failure      400    {object}  errorResponse
// @Failure      422    {object}  errorResponse
// @Router       /getFib [post]
func (s *Server) getFib(ctx *gin.Context) {
	var req getFibRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	fibSeq, err := s.service.GetFibSeq(ctx, req.Start, req.End)
	if err != nil {
		NewError(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, getFibResponse{Sequence: fibSeq})
}
