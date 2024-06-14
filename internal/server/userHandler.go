package server

import (
	// "database/sql"
	// "errors"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	db "github.com/ekefan/testBuildDeploy/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

// CreateUserRequest model for http post request body
type CreateUserRequest struct {
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Username    string `json:"username" binding:"required,alphanum"`
	Password    string `json:"password" binding:"required,min=8"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// UserResp model for http response body
type userResp struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// CustomUserResponse returns a UserResp object ready to be sent as a Response
func CustomUserResponse(user db.User) userResp {
	nameParts := strings.Fields(user.Fullname)
	// Assign values based on the length of nameParts
	var firstname, lastname string
	if len(nameParts) > 0 {
		firstname = nameParts[0]
	}
	if len(nameParts) > 1 {
		lastname = nameParts[1]
	}
	return userResp{
		Username:  user.Username,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     user.Email,
		// PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := utils.HashPassword(req.Password)
	/*
		if err != nil {
		  ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
	*/
	fullname := strings.Join([]string{req.Firstname, req.Lastname}, " ")
	args := db.CreateUserParams{
		Fullname:       fullname,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: req.Password,
		PhoneNumber:    req.PhoneNumber,
	}
	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		fmt.Println("not possible")
			if pqErr, ok := err.(*pq.Error); ok {
				fmt.Println(pqErr.Code.Name())
				switch pqErr.Code.Name() {
				case "unique_violation":
					ctx.JSON(http.StatusForbidden, errorResponse(err))
					return
				}
			}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := CustomUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	//AccessToken string `json:"access_token"`
	User userResp `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//check_password
	/*
		err = utils.CheckPassword(req.Password, user.HashedPassword)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		}

		accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	*/
	resp := loginUserResponse{
		//	AccessToken: accessToken,
		User: CustomUserResponse(user),
	}
	ctx.JSON(http.StatusOK, resp)
}
