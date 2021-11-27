package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) TestController() gin.HandlerFunc {
	type request struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password"`
	}

	return func(ctx *gin.Context) {
		req := &request{}
		if err := ctx.BindQuery(req); err != nil {
			ctx.JSON(400, gin.H{"success": false})
			return
		}
	}
}
