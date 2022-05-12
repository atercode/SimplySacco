package api

import (
	"database/sql"
	"net/http"

	db "github.com/atercode/SimplySacco/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createMemberRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	// StatusCode string `json:"status_code"`
}

func (server *Server) listMembers(ctx *gin.Context) {
	members, err := server.store.ListMembers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, members)
}

func (server *Server) createMember(ctx *gin.Context) {
	var req createMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateMemberParams{
		FullName:   req.FullName,
		Email:      req.Email,
		StatusCode: "TEST",
	}

	member, err := server.store.CreateMember(ctx.Request.Context(), args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, member)
}

type getMemberRequest struct {
	ID int32 `uri:"id" json:"id" binding:"required"`
}

func (server *Server) getMember(ctx *gin.Context) {
	var req getMemberRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	member, err := server.store.GetMember(ctx.Request.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, member)
}
