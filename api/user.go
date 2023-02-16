package api

import (
	"database/sql"
	db "github.com/VL-037/go-bank/db/sqlc"
	"github.com/VL-037/go-bank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username          string         `json:"username"`
	FullName          string         `json:"full_name"`
	Email             string         `json:"email"`
	PasswordUpdatedAt time.Time      `json:"password_updated_at"`
	CreatedBy         sql.NullString `json:"created_by"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedBy         sql.NullString `json:"updated_by"`
	UpdatedAt         time.Time      `json:"updated_at"`
	MarkForDelete     bool           `json:"mark_for_delete"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordUpdatedAt: user.PasswordUpdatedAt,
		CreatedBy:         user.CreatedBy,
		CreatedAt:         user.CreatedAt,
		UpdatedBy:         user.UpdatedBy,
		UpdatedAt:         user.UpdatedAt,
		MarkForDelete:     user.MarkForDelete,
	}

	ctx.JSON(http.StatusOK, response)
}
