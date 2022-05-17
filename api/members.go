package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/atercode/SimplySacco/db/sqlc"
	"github.com/atercode/SimplySacco/utils"
	"github.com/gin-gonic/gin"
)

type listMemberRequest struct {
	PageID     int32  `form:"page_id" binding:"required,min=1"`
	PageSize   int32  `form:"page_size" binding:"required,min=5,max=100"`
	StatusCode string `form:"status_code"`
}

type memberResponse struct {
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	StatusCode        string    `json:"status_code"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func singleMemberRepsonse(member db.Member) memberResponse {
	return memberResponse{
		FullName:          member.FullName,
		Email:             member.Email,
		StatusCode:        member.StatusCode,
		PasswordChangedAt: member.PasswordChangedAt,
		CreatedAt:         member.CreatedAt.Time,
	}
}

func (server *Server) listMembers(ctx *gin.Context) {
	var req listMemberRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// if len(req.StatusCode) > 0 {
	// 	//check if status code is valid
	// 	status, err := server.store.GetStatus(context.Background(), req.StatusCode)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 		return
	// 	}
	// 	status_code = status.Code
	// }
	args := db.ListMembersPaginatedParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	log.Print("Args are ", args)
	members, err := server.store.ListMembersPaginated(ctx.Request.Context(), args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var formatted_members = []memberResponse{}
	for _, member := range members {
		formatted_members = append(formatted_members, singleMemberRepsonse(member))
	}
	ctx.JSON(http.StatusOK, formatted_members)
}

type createMemberRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,alphanum,min=6"`
	// StatusCode string `json:"status_code"`
}

func (server *Server) createMember(ctx *gin.Context) {
	var req createMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.CreateMemberParams{
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		StatusCode:     "TEST",
	}

	member, err := server.store.CreateMember(ctx.Request.Context(), args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, singleMemberRepsonse(member))
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

type loginMemberRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginMemberResponse struct {
	AccessToken string         `json:access_token`
	Member      memberResponse `json:member`
}

func (server *Server) loginMember(ctx *gin.Context) {
	var req loginMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	member, err := server.store.GetMemberByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, member.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(member.Email, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginMemberResponse{
		AccessToken: accessToken,
		Member:      singleMemberRepsonse(member),
	}

	ctx.JSON(http.StatusOK, rsp)
}
